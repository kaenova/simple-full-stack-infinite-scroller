import Head from 'next/head'
import TextContainer from '../components/TextContainer'

export default function Home() {
  return (
    <div className="flex flex-col items-center justify-center min-h-screen py-2 bg-gray-200">
      <Head>
        <title>Create Next App</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="flex flex-col items-center justify-center w-full flex-1 px-20 text-center">
        <h1 className="text-6xl font-bold">
          Welcome to Infinite Scroll Testing
        </h1>

        <p className="mt-3 text-2xl">
          Just nyoba-nyoba
        </p>

        <div className="flex flex-wrap items-center justify-around max-w-4xl mt-6 sm:w-full">
          <TextContainer>Diubah lagi gan</TextContainer>
        </div>
      </main>
    </div>
  )
}
