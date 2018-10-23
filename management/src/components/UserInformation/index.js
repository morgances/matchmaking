import styles from './index.less';
import { Button, InputNumber } from 'antd';

export default (props) => {
  return (
    <div>
      <div className={styles.line}>
        <p className={styles.lineTitle}>ID：</p>
        <p className={styles.lineContent}>{props.item.ID}</p>
      </div>
      
      <div className={styles.line}>
        <p className={styles.lineTitle}>昵称：</p>
        <p className={styles.lineContent}>{props.item.NickName}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>真实姓名：</p>
        <p className={styles.lineContent}>{props.item.RealName}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>性别：</p>
        <p className={styles.lineContent}>{props.item.Sex}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>年龄：</p>
        <p className={styles.lineContent}>{props.item.Age}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>身高：</p>
        <p className={styles.lineContent}>{props.item.Height}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>地址：</p>
        <p className={styles.lineContent}>{props.item.Location}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>工作：</p>
        <p className={styles.lineContent}>{props.item.Job}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>信仰：</p>
        <p className={styles.lineContent}>{props.item.Faith}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>星座：</p>
        <p className={styles.lineContent}>{props.item.Constellation}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>自我介绍：</p>
        <p className={styles.lineContent}>{props.item.SelfIntroduction}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>择偶标准：</p>
        <p className={styles.lineContent}>{props.item.SelecCriteria}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>认证：</p>
        <p className={styles.lineContent}>
          <Button size="small" type="primary" ghost>{props.item.Certified}</Button>
          认证
        </p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>Vip：</p>
        <p className={styles.lineContent}>{props.item.Vip}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>积分：</p>
        <p className={styles.lineContent}>{props.item.Points}分</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>玫瑰：</p>
        <p className={styles.lineContent}>{props.item.Rose}朵</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>魅力值：</p>
        <p className={styles.lineContent}>{props.item.Charm}分</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>免费相亲机会：</p>
        <p className={styles.lineContent}>
          <InputNumber min={0} max={10} defaultValue={props.item.DatePrivilege}></InputNumber>
          次
        </p>
      </div>
    </div>
  );
}
