import { Component } from 'react';
import dva from 'dva';
import createLoading from 'dva-loading';

let app = dva({
  history: window.g_history,
  
});

window.g_app = app;
app.use(createLoading());

app.model({ namespace: 'global', ...(require('/Users/a11/Github/group-matchmaking/management/src/models/global.js').default) });
app.model({ namespace: 'list', ...(require('/Users/a11/Github/group-matchmaking/management/src/models/list.js').default) });
app.model({ namespace: 'login', ...(require('/Users/a11/Github/group-matchmaking/management/src/models/login.js').default) });
app.model({ namespace: 'project', ...(require('/Users/a11/Github/group-matchmaking/management/src/models/project.js').default) });
app.model({ namespace: 'setting', ...(require('/Users/a11/Github/group-matchmaking/management/src/models/setting.js').default) });
app.model({ namespace: 'user', ...(require('/Users/a11/Github/group-matchmaking/management/src/models/user.js').default) });
app.model({ namespace: 'activities', ...(require('/Users/a11/Github/group-matchmaking/management/src/pages/Dashboard/models/activities.js').default) });
app.model({ namespace: 'chart', ...(require('/Users/a11/Github/group-matchmaking/management/src/pages/Dashboard/models/chart.js').default) });
app.model({ namespace: 'monitor', ...(require('/Users/a11/Github/group-matchmaking/management/src/pages/Dashboard/models/monitor.js').default) });
app.model({ namespace: 'form', ...(require('/Users/a11/Github/group-matchmaking/management/src/pages/Forms/models/form.js').default) });
app.model({ namespace: 'rule', ...(require('/Users/a11/Github/group-matchmaking/management/src/pages/List/models/rule.js').default) });
app.model({ namespace: 'profile', ...(require('/Users/a11/Github/group-matchmaking/management/src/pages/Profile/models/profile.js').default) });
app.model({ namespace: 'error', ...(require('/Users/a11/Github/group-matchmaking/management/src/pages/Exception/models/error.js').default) });
app.model({ namespace: 'geographic', ...(require('/Users/a11/Github/group-matchmaking/management/src/pages/Account/Settings/models/geographic.js').default) });

class DvaContainer extends Component {
  render() {
    app.router(() => this.props.children);
    return app.start()();
  }
}

export default DvaContainer;
