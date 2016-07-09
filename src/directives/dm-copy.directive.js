(function () {
    'use strict';
    angular.module('glance')
        .directive('dmCopy', dmCopy);

    dmCopy.$inject = ["Notification"];

    function dmCopy(Notification) {
        return {
            restrict: 'A',
            scope: {
                target: '@dmCopy'
            },
            link: link
        };
        
        function link(scope, el, attrs) {
            el.attr('data-clipboard-target', scope.target);
            var clip = new ZeroClipboard(el, {
                moviePath: "/bower_components/zeroclipboard/dist/ZeroClipboard.swf"
            });
            clip.on("aftercopy", function (event) {
                Notification.success('复制成功');
            });
            clip.on("error", function () {
                el.click(function () {
                    Notification.warning('复制命令出现了问题，请您在页面选择命令并复制');
                });
            });
        }
    }
})();