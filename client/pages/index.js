import Head from 'next/head'
import styles from '../styles/Home.module.css'
import Homepage from './components/homepage'

export default function Home() {
  
  return (
    <div className={styles.container}>
      <Head>
        <title>MySpotify</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>
        <Homepage />
      </main>
    </div>
  )
}
