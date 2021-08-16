import User, {IUser} from '@entities/User';
import UserLogin, {IUserLogin} from '@entities/UserLogin';

class Auth {
    static async current() {
        return await fetch('/user', {method: 'GET'});
    }

    static async login(userlogin: IUserLogin) {
        const haveError: string[] = [];
        const userloginValid = new UserLogin(userlogin);
        haveError.push(...userloginValid.validate());
        if (haveError.length > 0) {
            return Promise.resolve(haveError);
        }
        return await fetch('/users/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(userloginValid.serialize()),
        });
    }

    static async save(user: IUser) {
        const haveError: string[] = [];
        const userValid = new User(user);
        haveError.push(...userValid.validate());
        if (haveError.length > 0) {
            return Promise.resolve(haveError);
        }
        return await fetch('/user', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(userValid.serialize()),
        });
    }
}
