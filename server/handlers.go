package server

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"interactive_course/runner"
)

type Server struct {
	ContentDir  string
	ProgressDir string
	mu          sync.Mutex
}

func New(contentDir, progressDir string) *Server {
	return &Server{
		ContentDir:  contentDir,
		ProgressDir: progressDir,
	}
}

type Lesson struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Exercises []Exercise `json:"exercises"`
}

type Exercise struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type SubmitRequest struct {
	Code       string `json:"code"`
	LessonID   string `json:"lessonId"`
	ExerciseID string `json:"exerciseId"`
}

type ProgressData struct {
	Completed map[string]bool `json:"completed"`
}

func (s *Server) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/lessons", s.handleLessons)
	mux.HandleFunc("/api/lesson/", s.handleLesson)
	mux.HandleFunc("/api/exercise/", s.handleExercise)
	mux.HandleFunc("/api/submit", s.handleSubmit)
	mux.HandleFunc("/api/progress", s.handleProgress)
}

func (s *Server) handleLessons(w http.ResponseWriter, r *http.Request) {
	lessons := s.loadLessons()
	writeJSON(w, lessons)
}

func (s *Server) handleLesson(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/api/lesson/"), "/")
	if len(parts) < 1 {
		http.Error(w, "missing lesson id", http.StatusBadRequest)
		return
	}
	lessonID := parts[0]

	mdPath := filepath.Join(s.ContentDir, "lessons", lessonID, "lesson.md")
	data, err := os.ReadFile(mdPath)
	if err != nil {
		http.Error(w, "lesson not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(data)
}

func (s *Server) handleExercise(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/api/exercise/"), "/")
	if len(parts) < 2 {
		http.Error(w, "missing lesson/exercise id", http.StatusBadRequest)
		return
	}
	lessonID := parts[0]
	exerciseID := parts[1]

	exerciseDir := filepath.Join(s.ContentDir, "lessons", lessonID, "exercises", exerciseID)

	exerciseMD, _ := os.ReadFile(filepath.Join(exerciseDir, "exercise.md"))
	starter, _ := os.ReadFile(filepath.Join(exerciseDir, "starter.go"))

	resp := map[string]string{
		"exercise": string(exerciseMD),
		"starter":  string(starter),
	}
	writeJSON(w, resp)
}

func (s *Server) handleSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SubmitRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	cfg := runner.DefaultConfig(s.ContentDir)

	lessonDir := filepath.Join(s.ContentDir, "lessons", req.LessonID)
	sortMeta := filepath.Join(lessonDir, "exercises", req.ExerciseID, "sort_output")
	if _, err := os.Stat(sortMeta); err == nil {
		cfg.SortOutput = true
	}

	result := runner.Run(req.Code, req.LessonID, req.ExerciseID, cfg)

	if result.Passed {
		s.markComplete(req.LessonID, req.ExerciseID)
	}

	writeJSON(w, result)
}

func (s *Server) handleProgress(w http.ResponseWriter, r *http.Request) {
	progress := s.loadProgress()
	writeJSON(w, progress)
}

func (s *Server) loadLessons() []Lesson {
	lessonsDir := filepath.Join(s.ContentDir, "lessons")
	entries, err := os.ReadDir(lessonsDir)
	if err != nil {
		return nil
	}

	progress := s.loadProgress()
	var lessons []Lesson

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		lessonID := entry.Name()
		title := extractTitle(filepath.Join(lessonsDir, lessonID, "lesson.md"))

		exercisesDir := filepath.Join(lessonsDir, lessonID, "exercises")
		exEntries, _ := os.ReadDir(exercisesDir)

		var exercises []Exercise
		for _, ex := range exEntries {
			if !ex.IsDir() {
				continue
			}
			exTitle := extractTitle(filepath.Join(exercisesDir, ex.Name(), "exercise.md"))
			key := lessonID + "/" + ex.Name()
			status := "pending"
			if progress.Completed[key] {
				status = "completed"
			}
			exercises = append(exercises, Exercise{
				ID:     ex.Name(),
				Title:  exTitle,
				Status: status,
			})
		}

		lessons = append(lessons, Lesson{
			ID:        lessonID,
			Title:     title,
			Exercises: exercises,
		})
	}

	sort.Slice(lessons, func(i, j int) bool {
		return lessons[i].ID < lessons[j].ID
	})

	return lessons
}

func (s *Server) loadProgress() ProgressData {
	s.mu.Lock()
	defer s.mu.Unlock()

	path := filepath.Join(s.ProgressDir, "progress.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return ProgressData{Completed: make(map[string]bool)}
	}

	var progress ProgressData
	if err := json.Unmarshal(data, &progress); err != nil {
		return ProgressData{Completed: make(map[string]bool)}
	}
	if progress.Completed == nil {
		progress.Completed = make(map[string]bool)
	}
	return progress
}

func (s *Server) markComplete(lessonID, exerciseID string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	path := filepath.Join(s.ProgressDir, "progress.json")
	var progress ProgressData

	data, err := os.ReadFile(path)
	if err == nil {
		json.Unmarshal(data, &progress)
	}
	if progress.Completed == nil {
		progress.Completed = make(map[string]bool)
	}

	key := lessonID + "/" + exerciseID
	progress.Completed[key] = true

	newData, _ := json.MarshalIndent(progress, "", "  ")
	os.WriteFile(path, newData, 0644)
}

func extractTitle(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return filepath.Base(filepath.Dir(path))
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "# ") {
			return strings.TrimPrefix(line, "# ")
		}
	}
	return filepath.Base(filepath.Dir(path))
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
