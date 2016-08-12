(function () {
    'use strict';
    angular.module('app.stack')
        .factory('stackBackend', stackBackend);


    /* @ngInject */
    function stackBackend(gHttp) {
        return {
            createStack: createStack,
            listStacks: listStacks,
            getStack: getStack,
            getService: getService,
            listStackServices: listStackServices,
            upServiceScale: upServiceScale,
            deleteStack: deleteStack,
            listServiceTasks: listServiceTasks,
            updateService: updateService,
            getServiceCDUrl: getServiceCDUrl
        };

        function createStack(data, form, groupId) {
            var params = {
                //TODO: disabled group function, so I am setting group_id = 1 bypass the func
                //group_id: groupId
                group_id: 1
            };
            return gHttp.Resource('stack.stacks').post(data, {form: form, params: params});
        }

        function listStacks() {
            return gHttp.Resource('stack.stacks').get();
        }

        function getStack(stackName) {
            return gHttp.Resource('stack.stack', {stack_name: stackName}).get();
        }

        function deleteStack(stackName) {
            return gHttp.Resource('stack.stack', {stack_name: stackName}).delete();
        }

        function listStackServices(stackName) {
            return gHttp.Resource('stack.services', {stack_name: stackName}).get();
        }

        function upServiceScale(stackName, serviceID, scale) {
            return gHttp.Resource('stack.service', {
                stack_name: stackName,
                service_id: serviceID
            }).patch({NumTasks: scale});
        }

        function getService(stackName, serviceID) {
            return gHttp.Resource('stack.service', {stack_name: stackName, service_id: serviceID}).get();
        }

        function getServiceCDUrl(stackName, serviceID) {
            return gHttp.Resource('stack.serviceCdUrl', {stack_name: stackName, service_id: serviceID}).get();
        }

        function listServiceTasks(stackName, serviceID) {
            return gHttp.Resource('stack.tasks', {stack_name: stackName, service_id: serviceID}).get();
        }

        function updateService(data, form, stackName, serviceID) {
            return gHttp.Resource('stack.service', {
                stack_name: stackName,
                service_id: serviceID
            }).put(data, {form: form});
        }
    }
})();
