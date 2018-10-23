import styles from './index.less';
import { Button, InputNumber } from 'antd';

export default (props) => {
  return (
    <div>
      <div className={styles.line}>
        <p className={styles.lineTitle}>ID：</p>
        <p className={styles.lineContent}>{props.item.open_id}</p>
      </div>
      
      <div className={styles.line}>
        <p className={styles.lineTitle}>昵称：</p>
        <p className={styles.lineContent}>{props.item.nick_name}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>真实姓名：</p>
        <p className={styles.lineContent}>{props.item.real_name}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>性别：</p>
        <p className={styles.lineContent}>{props.item.sex}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>年龄：</p>
        <p className={styles.lineContent}>{props.item.age}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>身高：</p>
        <p className={styles.lineContent}>{props.item.height}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>地址：</p>
        <p className={styles.lineContent}>{props.item.location}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>工作：</p>
        <p className={styles.lineContent}>{props.item.job}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>信仰：</p>
        <p className={styles.lineContent}>{props.item.faith}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>星座：</p>
        <p className={styles.lineContent}>{props.item.constellation}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>自我介绍：</p>
        <p className={styles.lineContent}>{props.item.self_introduction}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>择偶标准：</p>
        <p className={styles.lineContent}>{props.item.selec_criteria}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>认证：</p>
        <p className={styles.lineContent}>
          <Button size="small" type="primary" ghost>{props.item.certified}</Button>
          认证
        </p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>Vip：</p>
        <p className={styles.lineContent}>{props.item.vip}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>积分：</p>
        <p className={styles.lineContent}>{props.item.points}分</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>玫瑰：</p>
        <p className={styles.lineContent}>{props.item.rose}朵</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>魅力值：</p>
        <p className={styles.lineContent}>{props.item.charm}分</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>免费相亲机会：</p>
        <p className={styles.lineContent}>
          <InputNumber min={0} max={10} defaultValue={props.item.date_privilege}></InputNumber>
          次
        </p>
      </div>
    </div>
  );
}
