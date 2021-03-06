/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import '@fbcnms/babel-register/polyfill';

import LoginForm from '@fbcnms/ui/components/auth/LoginForm.js';
import React from 'react';
import ReactDOM from 'react-dom';
import fbt from 'fbt';
import nullthrows from '@fbcnms/util/nullthrows';
import {AppContextProvider} from '@fbcnms/ui/context/AppContext';
import {BrowserRouter} from 'react-router-dom';

import {} from './common/axiosConfig';
import {useRouter} from '@fbcnms/ui/hooks';

function LoginWrapper() {
  const {history, location} = useRouter();
  let error;
  if (location.search.includes('invalid=true')) {
    error = fbt(
      'Invalid login credentials',
      'Login error when invalid credentials are used',
    );
  }
  return (
    <LoginForm
      action={history.createHref({pathname: '/user/login'})}
      ssoAction={history.createHref({pathname: '/user/login/saml'})}
      title={fbt('Connectivity Platform', 'Main page title')}
      isSSO={window.CONFIG.appData.isSSO}
      csrfToken={window.CONFIG.appData.csrfToken}
      error={error}
    />
  );
}

ReactDOM.render(
  <AppContextProvider>
    <BrowserRouter basename="/">
      <LoginWrapper />
    </BrowserRouter>
  </AppContextProvider>,
  nullthrows(document.getElementById('root')),
);
