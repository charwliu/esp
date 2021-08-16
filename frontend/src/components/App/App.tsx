/**
 * Created by: Andrey Polyakov (andrey@polyakov.im)
 */
import ErrorBoundary from 'antd/lib/alert/ErrorBoundary';
import React, {Suspense, lazy} from 'react';
import {BrowserRouter, Route, Routes} from 'react-router-dom';

import Icons from '@components/Icons';

export const App = (): React.ReactElement => {
    return (
        <ErrorBoundary>
            <BrowserRouter>
                <Routes>
                    <Route path="/" element={<Home />} />
                    <Route path="dashboard" element={<Dashboard />} />
                    <Route path="*" element={<NotFound />} />
                </Routes>
                <Icons />
            </BrowserRouter>
        </ErrorBoundary>
    );
};

const Home = (): React.ReactElement => {
    return <div>Home</div>;
};

const NotFound = (): React.ReactElement => {
    return <div>NotFound</div>;
};

const Dashboard = (): React.ReactElement => {
    return <div>NotFound</div>;
};
