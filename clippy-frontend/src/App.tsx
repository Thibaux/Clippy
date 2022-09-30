import React from 'react';
import styles from './App.module.scss';
import { ItemsOverview } from './components/ItemsOverview/ItemsOverview';

const App = () => {
  return (
    <div className={styles.appWrapper}>
      <ItemsOverview />;
    </div>
  );
};

export default App;
