// Karma configuration
// Generated on Fri Jul 08 2016 19:49:56 GMT+0800 (CST)

module.exports = function (config) {
    config.set({

        // base path that will be used to resolve all patterns (eg. files, exclude)
        basePath: '',


        // frameworks to use
        // available frameworks: https://npmjs.org/browse/keyword/karma-adapter
        frameworks: ['jasmine'],


        // list of files / patterns to load in the browser
        files: [
            'bower_components/jquery/dist/jquery.js',
            'bower_components/bootstrap/dist/js/bootstrap.js',
            'bower_components/codemirror/lib/codemirror.js',
            'bower_components/codemirror/mode/yaml/yaml.js',
            'bower_components/angular/angular.js',
            'bower_components/angular-cookies/angular-cookies.js',
            'bower_components/angular-messages/angular-messages.js',
            'bower_components/angular-ui-router/release/angular-ui-router.js',
            'bower_components/angular-ui-codemirror/ui-codemirror.js',
            'bower_components/angular-sanitize/angular-sanitize.js',
            'bower_components/angular-aria/angular-aria.js',
            'bower_components/angular-material/angular-material.js',
            'bower_components/ngSocket/dist/ngSocket.js',
            'bower_components/zeroclipboard/dist/ZeroClipboard.js',
            'bower_components/angular-animate/angular-animate.min.js',
            'bower_components/angular-ui-notification/dist/angular-ui-notification.min.js',
            'bower_components/echarts/src/echarts.js',
            'bower_components/showdown/src/showdown.js',
            'bower_components/angular-markdown-directive/markdown.js',
            'bower_components/angular-base64/angular-base64.min.js',
            'bower_components/angular-material-data-table/dist/md-data-table.js',
            'bower_components/angular-loading-bar/build/loading-bar.js',
            'bower_components/angular-mocks/angular-mocks.js'
        ],


        // list of files to exclude
        exclude: [],


        // preprocess matching files before serving them to the browser
        // available preprocessors: https://npmjs.org/browse/keyword/karma-preprocessor
        preprocessors: {},


        // test results reporter to use
        // possible values: 'dots', 'progress'
        // available reporters: https://npmjs.org/browse/keyword/karma-reporter
        reporters: ['progress'],


        // web server port
        port: 9876,


        // enable / disable colors in the output (reporters and logs)
        colors: true,


        // level of logging
        // possible values: config.LOG_DISABLE || config.LOG_ERROR || config.LOG_WARN || config.LOG_INFO || config.LOG_DEBUG
        logLevel: config.LOG_INFO,


        // enable / disable watching file and executing tests whenever any file changes
        autoWatch: true,


        // start these browsers
        // available browser launchers: https://npmjs.org/browse/keyword/karma-launcher
        browsers: ['Chrome'],


        // Continuous Integration mode
        // if true, Karma captures browsers, runs the tests and exits
        singleRun: false,

        // Concurrency level
        // how many browser should be started simultaneous
        concurrency: Infinity
    })
}
