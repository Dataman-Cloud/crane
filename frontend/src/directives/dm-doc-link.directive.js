(function () {
    'use strict';
    angular.module('app')
        .directive('dmDocLink', dmDocLink);

    /* @ngInject */
    function dmDocLink() {
        return {
            restrict: 'A',
            link: link
        };

        function link(scope, elem, attrs) {
            var link = attrs['dmDocLink'] || attr['dm-doc-link'];
            var icon = "<i class='fa fa-umbrella'></i>";
            var alink = angular.element("<a href='" + link + "' target='_blank' class='help-info' style='position: relative; display: inline-block;'></a>");
            var text = elem.html()
            alink.append(icon + text)

            elem.html(alink)
        }
    }
})();
