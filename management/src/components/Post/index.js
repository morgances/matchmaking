import styles from './index.less';
import { Button, Tag, Avatar } from 'antd';

export default (props) => {
  return (
    <div className={styles.auditPost}>
      <div className={styles.post}>
        <div className={styles.CardID}>
          <Tag className={styles.tag} color="red">ID：{props.item.open_id}</Tag>
          <Avatar size={60} src={props.item.avatar} />
        </div>

        <div className={styles.cardContent}>
          <h1>{props.item.title}</h1>
          <p>{props.item.content}</p>
          <img className={styles.img} alt="甘霖娘" src={props.item.img} />
        </div>
      </div>

      <div className={styles.button}>
        <Button type="primary" size="large">可通过审核</Button>
        <br />
        <Button type="danger" size="large">不可通过审核</Button>
      </div>
    </div>
  );
}
