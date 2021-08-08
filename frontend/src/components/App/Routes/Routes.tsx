import React, { FC, useContext } from 'react';
import { Layout } from 'antd';
import { Switch, Route, Redirect } from 'react-router-dom';
import { observer } from 'mobx-react-lite';

import Store from 'Store';
import { Paths } from '../../common/Paths';

import { Login, ForgotPassword } from '../Login';

import s from './Routes.module.pcss';
import Sidebar from '../Sidebar';
import Dashboard from '../Dashboard';
import Header from '../Header';
import SetupGuide from '../SetupGuide';
import { GeneralSettings } from '../Settings';

const {Content} = Layout

const AuthRoutes: FC = React.memo(() => {
    return (
        <Switch>
            <Route
                exact
                path={Paths.ForgotPassword}
                component={ForgotPassword}
            />
            <Route
                path={Paths.Login}
                component={Login}
            />
        </Switch>
    );
});


const AppRoutes: FC = observer(() => {
    return (
        <Layout className={s.app}>
            <Sidebar />
            <Layout>
                <Header />
                <Content>
                    <Switch>
                        <Route
                            exact
                            path={Paths.Dashboard}
                            component={Dashboard}
                        />
                        <Route
                            exact
                            path={Paths.SetupGuide}
                            component={SetupGuide}
                        />
                        <Route
                            exact
                            path={Paths.SettingsGeneral}
                            component={GeneralSettings}
                        />
                        <Redirect to={Paths.Dashboard} />
                    </Switch>
                </Content>
            </Layout>
        </Layout>
    );
});

const Routes: FC = observer(() => {
    const store = useContext(Store);
    const { login: { loggedIn } } = store;
    if (loggedIn) {
        return <AppRoutes />;
    }
    return <AuthRoutes />;
});

export default Routes;

