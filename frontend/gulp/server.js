var gulp = require('gulp');
var connect = require('gulp-connect');
port = process.env.port || 5000;

//live reload
gulp.task('connect', function () {
    connect.server({
        root: './',
        port: port,
        livereload: true,
        fallback: 'index.html'
    })
});

//reload js
gulp.task('js', function () {
    gulp.src('src/**/*.js')
        .pipe(connect.reload())
});

//reload html
gulp.task('html', function () {
    gulp.src('src/**/*.html')
        .pipe(connect.reload())
});

//watch
gulp.task('watch', function () {
    gulp.watch('src/**/*.js', ['js']);
    gulp.watch('src/**/*.html', ['html']);
});

gulp.task('serve', ['connect', 'watch']);