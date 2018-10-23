import styles from './index.less';
import { Button, InputNumber } from 'antd';

export default (props) => {
  return (
    <div className={styles.auditPost}>
      <div className={styles.post}>
        <div className={styles.CardID}>
          <Tag className={styles.tag} color="red">ID：123456789</Tag>
          <Avatar size={60} src="https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png"/>
        </div>

        <div className={styles.cardContent}>
          <h1>同城同性约</h1>
          <p>我是大四的基佬，喜欢搞基，有意者私聊，猛戳→→微信号：技术猫</p>
          <img className={styles.img} alt="甘霖娘" src="https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png" />
        </div>
      </div>

      <div>
        <Button type="primary" size="large">可通过审核</Button>
        <br />
        <br />
        <Button type="danger" size="large">不可通过审核</Button>
      </div>
    </div>
  );
}
