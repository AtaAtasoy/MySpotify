import Head from 'next/head'
import styles from '../styles/Home.module.css'
import LoginButton from './components/login-btn'

export default function Home() {
  return (
    <div className={styles.container}>
      <Head>
        <title>MySpotify</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>
          <LoginButton/>
      </main>
    </div>
  )
}
