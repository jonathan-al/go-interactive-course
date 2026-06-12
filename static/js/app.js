let editor;
let lessons = [];
let currentLesson = null;
let currentExercise = null;
let allExercises = [];

document.addEventListener('DOMContentLoaded', () => {
    editor = CodeMirror.fromTextArea(document.getElementById('codeEditor'), {
        mode: 'go',
        theme: 'dracula',
        lineNumbers: true,
        autoCloseBrackets: true,
        matchBrackets: true,
        indentUnit: 4,
        tabSize: 4,
        indentWithTabs: true,
    });

    loadLessons();

    document.getElementById('runBtn').addEventListener('click', runCode);
    document.getElementById('resetBtn').addEventListener('click', resetCode);
    document.getElementById('prevBtn').addEventListener('click', prevExercise);
    document.getElementById('nextBtn').addEventListener('click', nextExercise);
    document.getElementById('menuToggle').addEventListener('click', toggleSidebar);
    document.getElementById('exerciseSelect').addEventListener('change', (e) => {
        loadExercise(currentLesson, e.target.value);
    });
});

async function loadLessons() {
    const res = await fetch('/api/lessons');
    lessons = await res.json();
    renderSidebar();
    updateProgress();
    buildExerciseIndex();
}

function buildExerciseIndex() {
    allExercises = [];
    lessons.forEach(lesson => {
        lesson.exercises.forEach(ex => {
            allExercises.push({ lessonId: lesson.id, exerciseId: ex.id, title: ex.title, status: ex.status });
        });
    });
}

function renderSidebar() {
    const nav = document.getElementById('lessonsNav');
    nav.innerHTML = '';

    lessons.forEach(lesson => {
        const group = document.createElement('div');
        group.className = 'lesson-group';

        const title = document.createElement('div');
        title.className = 'lesson-group-title';
        title.innerHTML = `<span class="arrow">&#9654;</span> ${lesson.title}`;

        const list = document.createElement('div');
        list.className = 'exercise-list';

        lesson.exercises.forEach(ex => {
            const item = document.createElement('div');
            item.className = 'exercise-item';
            item.dataset.lesson = lesson.id;
            item.dataset.exercise = ex.id;
            const check = ex.status === 'completed' ? '<span class="check">&#10003;</span>' : '<span class="check">&nbsp;&nbsp;</span>';
            item.innerHTML = `${check} ${ex.title}`;
            item.addEventListener('click', () => {
                selectExercise(lesson.id, ex.id);
            });
            list.appendChild(item);
        });

        title.addEventListener('click', () => {
            title.classList.toggle('expanded');
            list.classList.toggle('show');
        });

        group.appendChild(title);
        group.appendChild(list);
        nav.appendChild(group);
    });
}

function updateProgress() {
    let total = 0;
    let completed = 0;
    lessons.forEach(l => {
        l.exercises.forEach(ex => {
            total++;
            if (ex.status === 'completed') completed++;
        });
    });
    const pct = total > 0 ? Math.round((completed / total) * 100) : 0;
    document.getElementById('progressFill').style.width = pct + '%';
    document.getElementById('progressText').textContent = `${completed}/${total} exercises (${pct}%)`;
}

function selectExercise(lessonId, exerciseId) {
    document.querySelectorAll('.exercise-item').forEach(el => el.classList.remove('active'));
    const target = document.querySelector(`.exercise-item[data-lesson="${lessonId}"][data-exercise="${exerciseId}"]`);
    if (target) target.classList.add('active');

    loadExercise(lessonId, exerciseId);
}

async function loadExercise(lessonId, exerciseId) {
    currentLesson = lessonId;
    currentExercise = exerciseId;

    const [lessonRes, exRes] = await Promise.all([
        fetch(`/api/lesson/${lessonId}`),
        fetch(`/api/exercise/${lessonId}/${exerciseId}`)
    ]);

    const lessonMD = await lessonRes.text();
    const exData = await exRes.json();

    document.getElementById('lessonContent').innerHTML = marked.parse(lessonMD);
    document.getElementById('exerciseDescription').innerHTML = marked.parse(exData.exercise);
    document.getElementById('currentTitle').textContent = `${lessonId} / ${exerciseId}`;

    editor.setValue(exData.starter);

    document.getElementById('exerciseSection').style.display = 'flex';
    document.getElementById('outputSection').style.display = 'none';
    document.getElementById('errorSection').style.display = 'none';
    document.getElementById('status').textContent = '';
    document.getElementById('status').className = 'status';

    populateExerciseSelect(lessonId, exerciseId);
    updateNavButtons();

    setTimeout(() => editor.refresh(), 10);
}

function populateExerciseSelect(lessonId, exerciseId) {
    const select = document.getElementById('exerciseSelect');
    select.innerHTML = '';
    const lesson = lessons.find(l => l.id === lessonId);
    if (!lesson) return;

    lesson.exercises.forEach(ex => {
        const opt = document.createElement('option');
        opt.value = ex.id;
        opt.textContent = ex.title;
        if (ex.id === exerciseId) opt.selected = true;
        select.appendChild(opt);
    });
}

function updateNavButtons() {
    const idx = allExercises.findIndex(e => e.lessonId === currentLesson && e.exerciseId === currentExercise);
    document.getElementById('prevBtn').disabled = idx <= 0;
    document.getElementById('nextBtn').disabled = idx >= allExercises.length - 1;
}

async function runCode() {
    const btn = document.getElementById('runBtn');
    const status = document.getElementById('status');
    btn.disabled = true;
    status.textContent = 'Running...';
    status.className = 'status running';
    document.getElementById('outputSection').style.display = 'none';
    document.getElementById('errorSection').style.display = 'none';

    try {
        const res = await fetch('/api/submit', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                code: editor.getValue(),
                lessonId: currentLesson,
                exerciseId: currentExercise
            })
        });

        const result = await res.json();

        if (result.error) {
            document.getElementById('errorOutput').textContent = result.error;
            document.getElementById('errorSection').style.display = 'block';
            status.textContent = 'Compilation/Runtime Error';
            status.className = 'status failed';
        } else if (result.passed) {
            document.getElementById('actualOutput').textContent = result.output || '(no output)';
            document.getElementById('expectedOutput').textContent = result.expected;
            document.getElementById('outputSection').style.display = 'flex';
            status.textContent = 'Passed!';
            status.className = 'status passed';
            markExerciseComplete();
        } else {
            document.getElementById('actualOutput').textContent = result.output || '(no output)';
            document.getElementById('expectedOutput').textContent = result.expected;
            document.getElementById('outputSection').style.display = 'flex';
            status.textContent = 'Output does not match expected';
            status.className = 'status failed';
        }
    } catch (err) {
        document.getElementById('errorOutput').textContent = err.message;
        document.getElementById('errorSection').style.display = 'block';
        status.textContent = 'Request failed';
        status.className = 'status failed';
    }

    btn.disabled = false;
}

async function markExerciseComplete() {
    await loadLessons();
    const target = document.querySelector(`.exercise-item[data-lesson="${currentLesson}"][data-exercise="${currentExercise}"]`);
    if (target) target.classList.add('active');
}

function resetCode() {
    loadExercise(currentLesson, currentExercise);
}

function prevExercise() {
    const idx = allExercises.findIndex(e => e.lessonId === currentLesson && e.exerciseId === currentExercise);
    if (idx > 0) {
        const prev = allExercises[idx - 1];
        selectExercise(prev.lessonId, prev.exerciseId);
    }
}

function nextExercise() {
    const idx = allExercises.findIndex(e => e.lessonId === currentLesson && e.exerciseId === currentExercise);
    if (idx < allExercises.length - 1) {
        const next = allExercises[idx + 1];
        selectExercise(next.lessonId, next.exerciseId);
    }
}

function toggleSidebar() {
    document.getElementById('sidebar').classList.toggle('open');
}
