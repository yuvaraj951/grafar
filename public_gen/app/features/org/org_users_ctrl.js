///<reference path="../../headers/common.d.ts" />
System.register(['../../core/core_module'], function(exports_1) {
    var core_module_1;
    var OrgUsersCtrl;
    return {
        setters:[
            function (core_module_1_1) {
                core_module_1 = core_module_1_1;
            }],
        execute: function() {
            OrgUsersCtrl = (function () {
                /** @ngInject */
                function OrgUsersCtrl($scope, $http, backendSrv) {
                    this.$scope = $scope;
                    this.$http = $http;
                    this.backendSrv = backendSrv;
                    this.user = {
                        loginOrEmail: '',
                        role: 'Viewer',

                    };

                    this.get();
                    this.editor = { index: 0 };

                }
                OrgUsersCtrl.prototype.get = function () {
                    var _this = this;
                    this.backendSrv.get('/api/org/users')
                        .then(function (users) {
                        _this.users = users;
                    });
                    this.backendSrv.get('/api/org/invites')
                        .then(function (pendingInvites) {
                        _this.pendingInvites = pendingInvites;
                    });
                    this.backendSrv.get('/api/org/process')
                         .then(function (process) {
                          _this.process = process;
                                        });
                    this.backendSrv.get('/api/org/machine')
                                             .then(function (machines) {
                                              _this.machines = machines;
                                                            });
                };
                OrgUsersCtrl.prototype.updateOrgUser = function (user) {
                    this.backendSrv.patch('/api/org/users/' + user.userId, user);
                };

                OrgUsersCtrl.prototype.removeUser = function (user) {
                    var _this = this;
                    this.$scope.appEvent('confirm-modal', {
                        title: 'Delete',
                        text: 'Are you sure you want to delete user ' + user.login + '?',
                        yesText: "Delete",
                        icon: "fa-warning",
                        onConfirm: function () {
                            _this.removeUserConfirmed(user);
                        }
                    });
                };
                OrgUsersCtrl.prototype.removeUserConfirmed = function (user) {
                    var _this = this;
                    this.backendSrv.delete('/api/org/users/' + user.userId)
                        .then(function () {
                        _this.get();
                    });
                };


                //delete the process
                 OrgUsersCtrl.prototype.removeProcess = function (p) {
                                    var _this = this;
                                    this.$scope.appEvent('confirm-modal', {
                                        title: 'Delete',
                                        text: 'Are you sure you want to delete user ' + p.processName + '?',
                                        yesText: "Delete",
                                        icon: "fa-warning",
                                        onConfirm: function () {
                                            _this.removeProcessConfirmed(p);
                                        }
                                    });
                                };
                                OrgUsersCtrl.prototype.removeProcessConfirmed = function (p) {
                                    var _this = this;
                                    this.backendSrv.delete('/api/org/process/' + p.processId)
                                        .then(function () {
                                        _this.get();
                                    });
                                };

             // delete the machine
              OrgUsersCtrl.prototype.removeMachine = function (machine) {
                                                 var _this = this;
                                                 this.$scope.appEvent('confirm-modal', {
                                                     title: 'Delete',
                                                     text: 'Are you sure you want to delete Machine ' + machine.machineName + '?',
                                                     yesText: "Delete",
                                                     icon: "fa-warning",
                                                     onConfirm: function () {
                                                         _this.removeProcessConfirmed(machine);
                                                     }
                                                 });
                                             };
                                             OrgUsersCtrl.prototype.removeProcessConfirmed = function (machine) {
                                                 var _this = this;
                                                 this.backendSrv.delete('/api/org/machine/' + machine.machineId)
                                                     .then(function () {
                                                     _this.get();
                                                 });
                                             };

                OrgUsersCtrl.prototype.revokeInvite = function (invite, evt) {
                    var _this = this;
                    evt.stopPropagation();
                    this.backendSrv.patch('/api/org/invites/' + invite.code + '/revoke')
                        .then(function () {
                        _this.get();
                    });
                };
                OrgUsersCtrl.prototype.copyInviteToClipboard = function (evt) {
                    evt.stopPropagation();
                };
                OrgUsersCtrl.prototype.openInviteModal = function () {
                    var modalScope = this.$scope.$new();
                    modalScope.invitesSent = function () {
                        this.get();
                    };
                    this.$scope.appEvent('show-modal', {
                        src: 'public/app/features/org/partials/invite.html',
                        modalClass: 'invite-modal',
                        scope: modalScope
                    });
                };
                OrgUsersCtrl.prototype.openProcessModal = function () {
                                    var modalScope = this.$scope.$new();
                                    modalScope.processSent = function () {
                                        this.get();
                                    };
                                    this.$scope.appEvent('show-modal', {
                                        src: 'public/app/features/org/partials/addProcess.html',
                                        modalClass: 'invite-modal',
                                        scope: modalScope

                                    });
                                };
             OrgUsersCtrl.prototype.openMachineModal = function () {
                                                 var modalScope = this.$scope.$new();
                                                 modalScope.processSent = function () {
                                                     this.get();
                                                 };
                                                 this.$scope.appEvent('show-modal', {
                                                     src: 'public/app/features/org/partials/addMachine.html',
                                                     modalClass: 'invite-modal',
                                                     scope: modalScope

                                                 });
                                             };
                 OrgUsersCtrl.prototype.updateProcessModal = function (p) {
                     var modalScope = this.$scope.$new();

                      modalScope.processUpdateSent = function () {
                     this.get();
                    };
                   this.$scope.appEvent('show-modal', {
                     src: 'public/app/features/org/partials/updateProcess.html',
                      modalClass: 'invite-modal',
                     scope: modalScope
                     });
                     };
                return OrgUsersCtrl;
            })();
            exports_1("OrgUsersCtrl", OrgUsersCtrl);
            core_module_1.default.controller('OrgUsersCtrl', OrgUsersCtrl);
        }
    }
});
//# sourceMappingURL=org_users_ctrl.js.map
