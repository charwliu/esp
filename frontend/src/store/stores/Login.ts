import {IUser} from '@entities/User';
import {errorChecker} from '@helpers/apiErrors';
import {flow, makeAutoObservable, observable} from 'mobx';

export default class Login {
    rootStore: Store;

    loggedIn = false;

    constructor(rootStore: Store) {
        this.rootStore = rootStore;
        makeAutoObservable(this, {
            loggedIn: observable,
            rootStore: false,
            checkLoggedIn: flow,
            login: flow,
        });
        this.checkLoggedIn();
    }

    *checkLoggedIn() {
        // const response = yield authApi.
        this.loggedIn = true;
    }

    *login(login: IUser) {
        this.loggedIn = true;
        return true;
    }
}
