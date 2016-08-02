var gulp = require('gulp');
var Server = require('karma').Server;
var $ = require('gulp-load-plugins')();

//js code review
gulp.task('jsLint', function () {
    gulp.src('src/**/*.js')
        .pipe($.jshint())
        .pipe($.jshint.reporter());
});

gulp.task('test', ['jsLint'], function (done) {
    new Server({
        configFile: process.cwd() + '/karma.config.js',
        singleRun: true
    }, done).start();
});