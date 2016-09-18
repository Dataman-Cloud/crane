(function () {
    'use strict';
    angular.module('app.utils')
        .config(i18nEn);

    /* @ngInject */
    function i18nEn($translateProvider) {
        $translateProvider.translations('en-us', {
            //general
            "Create Time": 'Create Time',
            "Update Time": 'Update Time',
            "Operation": 'Operation',
            "Update": 'Update',

            //index module
            "LOGOUT": 'Logout',
            "Stack": 'Project',
            "Image": 'Image',
            "Facility": 'Facility',
            "Node": 'Node',
            "Network": 'Network',
            "Warehouse certification": 'Warehouse certification',
            "Information": 'Information',
            "Activation": 'Activation',

            //stack module
            "Stack List": 'Stack List',
            "Create Project": 'Create Project',
            "Service Update": 'Service Update',
            "Project Detail": 'Project Detail',
            "Service Detail": 'Service Detail',
            "Choose the way to create": 'Choose the way to create',
            "Create of DAB": 'DAB',
            "Create of Form": 'Guide',
            "Create of Shortcut": 'Shortcut',
            "Delete Project": 'Delete Project',
            "Add to template": 'Add to template',
            "Service List": 'Service List',
            "Service Name": 'Service Name',
            "Task Number": 'Task Number',
            "Running/Total": 'Running/Total',
            "Service Scale": 'Service Scale',
            "Modal title for delete stack": 'The project will not be able to recover after deletion, confirm to be deleted?',
        });
    }
})();
