/**
 * Created by: Andrey Polyakov (andrey@polyakov.im)
 */
import '@styles/styles.less';
import '@styles/styles.scss';

import Store, {storeValue} from '@store/index';
import React from 'react';
import ReactDom from 'react-dom';

import {App} from '@components/app/app';

const Container = () => {
    return (
        <Store.Provider value={storeValue}>
            <App />
        </Store.Provider>
    );
};

ReactDom.render(<App />, document.getElementById('root'));
