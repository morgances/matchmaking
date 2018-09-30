import React from 'react';
import { Router, Route, Switch } from 'dva/router';
import TestPage from './routes/TestPage';
import User from './routes/UserLayout';
import Basic from './routes/BasicLayout';

function RouterConfig({ history }) {
  return (
    <Router history={history}>
      <Switch>
        <Route path="/test" exact component={TestPage} />

        <Route path="/" exact component={User} />

        <Route path="/basic" exact component={Basic} />
      </Switch>
    </Router>
  );
}

export default RouterConfig;
