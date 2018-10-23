import React, { Component } from 'react';
import { Tag, Button, Avatar, Modal } from 'antd';

import styles from './DynamicPost.less';

const confirm = Modal.confirm;

class DynamicPost extends Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  showConfirm() {
    confirm({
      title: 'Are you sure delete this task?',
      content: 'Some descriptions',
      okText: 'Yes',
      okType: 'danger',
      cancelText: 'No',
      onOk() {
        console.log('OK');
      },
      onCancel() {
        console.log('Cancel');
      },
    });
  }

  render() {
    return (
      <div className={styles.auditPost}>
        <div className={styles.post}>
          <div className={styles.CardID}>
            <Tag className={styles.tag} color="red">ID：123456789</Tag>
            <Avatar size={60} src="https://is4-ssl.mzstatic.com/image/thumb/Purple117/v4/f9/87/e3/f987e3f2-edb6-496c-d14a-a3e5fa266531/mzl.hifizenz.png/246x0w.jpg"/>
          </div>

          <div className={styles.cardContent}>
            <h1>同城同性约</h1>
            <p>我是大四的基佬，喜欢搞基，有意者私聊，猛戳→→微信号：技术猫</p>
            <img className={styles.img} alt="甘霖娘" src="https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png" />
          </div>
        </div>

        <div>
          <Button
            type="primary"
            size="large"
            // onClick={() => showConfirm()}
          >
            可通过审核
          </Button>
          <br />
          <br />
          <Button type="danger" size="large">不可通过审核</Button>
        </div>
      </div>
    );
  }
}

export default DynamicPost;
