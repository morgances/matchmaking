import React, { Fragment } from 'react';
import Link from 'umi/link';
import { Icon } from 'antd';
import GlobalFooter from '@/components/GlobalFooter';
import styles from './UserLayout.less';

const copyright = (
  <Fragment>
    Copyright <Icon type="copyright" /> 2018 老子是刘琦，我的出品
  </Fragment>
);

class UserLayout extends React.PureComponent {
  render() {
    const { children } = this.props;
    return (
      <div className={styles.container}>
        <div className={styles.content}>
          <div className={styles.top}>
            <div className={styles.header}>
              <Link to="/">
                <span className={styles.title}>管理后台</span>
              </Link>
            </div>
            <div className={styles.desc}>以婚至上小程序后台管理系统</div>
          </div>
          {children}
        </div>

        <GlobalFooter
          copyright={copyright}
        />
      </div>
    );
  }
}

export default UserLayout;
