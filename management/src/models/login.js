import { routerRedux } from 'dva/router';
import { accountLogin } from '@/services/api';
import { getPageQuery } from '@/utils/utils';

export default {
  namespace: 'login',

  state: {
    status: undefined,
  },

  effects: {
    *login({ payload }, { call, put }) {
      const response = yield call(accountLogin, payload);
      console.log('打印出我输入的信息：', payload);
      yield put({
        type: 'changeLoginStatus',
        payload: response,
      });
      // Login successfully
      if (response === true) {
        const urlParams = new URL(window.location.href);
        const params = getPageQuery();
        let { redirect } = params;
        if (redirect) {
          const redirectUrlParams = new URL(redirect);
          if (redirectUrlParams.origin === urlParams.origin) {
            redirect = redirect.substr(urlParams.origin.length);
            if (redirect.startsWith('/#')) {
              redirect = redirect.substr(2);
            }
          } else {
            window.location.href = "/userpage";
            return;
          }
        }
        yield put(routerRedux.replace(redirect || '/'));
      }
    },
  },

  reducers: {
    changeLoginStatus(state, { payload }) {
      return {
        ...state,
        status: payload.status,
        type: payload.type,
      };
    },
  },
};
