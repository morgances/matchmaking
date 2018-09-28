import React from 'react';
import styles from './UserLayout.less';

class UserLayout extends React.PureComponent {
  // getPageTitle() {
  //   const { routerData, location } = this.props;
  //   const { pathname } = location;
  //   let title = 'Ant Design Pro';
  //   if (routerData[pathname] && routerData[pathname].name) {
  //     title = `${routerData[pathname].name} - Ant Design Pro`;
  //   }
  //   return title;
  // }

  render() {
    return (
      <div className={styles.container}>
        <div className={styles.desc}>初始化后台管理登录界面</div>
      </div>
    );
  }
}

export default UserLayout;
