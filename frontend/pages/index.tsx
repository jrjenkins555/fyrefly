import Head from 'next/head';
import Layout from '../components/layout';
import utilStyles from '../styles/utils.module.css'
import InputFileButton from '../components/ui/input-file-button';

export default function Home() {
  return (
    // simple Layout component
    <Layout>
      <InputFileButton></InputFileButton>
    </Layout>
  );
}