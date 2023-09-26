import Head from 'next/head';
import Layout, { siteTitle } from '../components/layout';
import FileUploadButton from '../components/FileUploadButton';
import utilStyles from '../styles/utils.module.css'

export default function Home() {
  return (
    // simple Layout component
    <Layout>
      <FileUploadButton></FileUploadButton>
    </Layout>
  );
}