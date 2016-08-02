var path = require('path');
var gulp = require('gulp');
var conf = require('./conf');

var $ = require('gulp-load-plugins')({
    pattern: ['gulp-*', 'main-bower-files', 'del']
});

gulp.task('copy-conf', function () {
    gulp.src('src/conf.js')
        .pipe(gulp.dest('dist/js/'));
});

gulp.task('copy-pics', ['copy-conf'], function () {
    gulp.src('pics/*')
        .pipe(gulp.dest('dist/pics/'));
});

gulp.task('clean', function () {
    return $.del([path.join(conf.paths.dist, '/'), path.join(conf.paths.tmp, '/')]);
});

gulp.task('build', ['copy-pics']);