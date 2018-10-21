import styles from './index.less';

export default (props) => {
  return (
    <div>
      <div className={styles.line}>
        <p className={styles.lineTitle}>{props.item.title}</p>
        <p className={styles.lineContent}>{props.item.content}</p>
      </div>
    </div>
  );
}
