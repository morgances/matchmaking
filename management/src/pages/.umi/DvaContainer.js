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
app.model({ namespace: 'post', ...(require('/Users/a11/Github/group-matchmaking/management/src/models/post.js').default) });
app.model({ namespace: 'project', ...(require('/Users/a11/Github/group-matchmaking/management/src/models/project.js').default) });
app.model({ namespace: 'setting', ...(require('/Users/a11/Github/group-matchmaking/management/src/models/setting.js').default) });
app.model({ namespace: 'user', ...(require('/Users/a11/Github/group-matchmaking/management/src/models/user.js').default) });
app.model({ namespace: 'userinformation', ...(require('/Users/a11/Github/group-matchmaking/management/src/models/userinformation.js').default) });
app.model({ namespace: 'rule', ...(require('/Users/a11/Github/group-matchmaking/management/src/pages/List/models/rule.js').default) });
app.model({ namespace: 'profile', ...(require('/Users/a11/Github/group-matchmaking/management/src/pages/Profile/models/profile.js').default) });

class DvaContainer extends Component {
  render() {
    app.router(() => this.props.children);
    return app.start()();
  }
}

export default DvaContainer;
