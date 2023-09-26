import Head from 'next/head';
import Image from 'next/image';
import styles from './layout.module.css';
import Link from 'next/link';

const siteTitle = 'AI Powered Survey Automation';

export default function Layout({ children }) {
  return (
    <div className={styles.container}>
      <div className={styles.title}>
        &#9889; {siteTitle} &#9889;
      </div>
      <main>{children}</main>
    </div>
  );
}