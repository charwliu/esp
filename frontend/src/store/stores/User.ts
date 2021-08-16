import AuthApi from '@apis/auth';
import UsersApi from '@apis/Users';
import {IUser} from '@entities/User';
import {action, flow, makeAutoObservable, observable} from 'mobx';

import {Store} from '../store';

class UserStore {
    rootStore: Store;

    loadingUser = false;
    currentUser: IUser;
    updatingUser = false;
    updatingUserErrors: Error;
    constructor(rootStore: Store) {
        this.rootStore = rootStore;
        makeAutoObservable(this, {
            loadingUser: observable,
            updatingUser: observable,
            updatingUserErrors: observable,
            rootStore: false,
            pullUser: flow,
            login: flow,
        });
        this.checkLoggedIn();
    }

    *pullUser() {
        this.loadingUser = true;
        const response = fetch(`/api/v1/user`, {
            method: 'GET',
        }).then;
    }
}
