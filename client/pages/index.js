import Head from 'next/head'
import styles from '../styles/Home.module.css'
import Homepage from './views/Homepage'

export default function Home() {
  
  return (
    <div className={styles.container}>
      <Head>
        <meta http-equiv="Content-Security-Policy" content="upgrade-insecure-requests" />
        <title>MySpotify</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>
        <Homepage />
      </main>
    </div>
  )
}
