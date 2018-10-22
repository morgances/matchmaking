import styles from './index.less';

export default (props) => {
  return (
    <div>
      <div className={styles.line}>
        <p className={styles.lineTitle}>ID：</p>
        <p className={styles.lineContent}>{props.item.ID}</p>
      </div>
      
      <div className={styles.line}>
        <p className={styles.lineTitle}>NickName：</p>
        <p className={styles.lineContent}>{props.item.NickName}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>RealName：</p>
        <p className={styles.lineContent}>{props.item.RealName}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>Sex：</p>
        <p className={styles.lineContent}>{props.item.Sex}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>Age：</p>
        <p className={styles.lineContent}>{props.item.Age}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>Height：</p>
        <p className={styles.lineContent}>{props.item.Height}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>Location：</p>
        <p className={styles.lineContent}>{props.item.Location}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>Job：</p>
        <p className={styles.lineContent}>{props.item.Job}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>Faith：</p>
        <p className={styles.lineContent}>{props.item.Faith}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>Constellation：</p>
        <p className={styles.lineContent}>{props.item.Constellation}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>SelfIntroduction：</p>
        <p className={styles.lineContent}>{props.item.SelfIntroduction}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>SelecCriteria：</p>
        <p className={styles.lineContent}>{props.item.SelecCriteria}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>Certified：</p>
        <p className={styles.lineContent}>{props.item.Certified}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>Vip：</p>
        <p className={styles.lineContent}>{props.item.Vip}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>Points：</p>
        <p className={styles.lineContent}>{props.item.Points}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>Rose：</p>
        <p className={styles.lineContent}>{props.item.Rose}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>Charm：</p>
        <p className={styles.lineContent}>{props.item.Charm}</p>
      </div>

      <div className={styles.line}>
        <p className={styles.lineTitle}>DatePrivilege：</p>
        <p className={styles.lineContent}>{props.item.DatePrivilege}</p>
      </div>
    </div>
  );
}
