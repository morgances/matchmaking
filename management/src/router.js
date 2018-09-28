import React from 'react';
import { Router, Route, Switch } from 'dva/router';
import IndexPage from './routes/IndexPage';
import Products from './routes/Products';
import UserLayout from './routes/UserLayout';

function RouterConfig({ history }) {
  return (
    <Router history={history}>
      <Switch>
        <Route path="/IndexPage" exact component={IndexPage} />

        <Route path="/products" exact component={Products} />

        <Route path="/" exact component={UserLayout} />
      </Switch>
    </Router>
  );
}

export default RouterConfig;
