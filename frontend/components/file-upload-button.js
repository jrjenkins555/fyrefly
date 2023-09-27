import React, { useState } from 'react';
import buttonStyles from './file-upload-button.module.css';
import Cookies from "js-cookie"

export default function FileUploadButton() {
  const [file, setFile] = useState(null);
  const [uploading, setUploading] = useState(false);

  const handleFileChange = (event) => {
    const selectedFile = event.target.files[0];
    setFile(selectedFile);
  };

  const handleUpload = async () => {
    if (!file) {
      alert('Please select a file to upload.');
      return;
    }

    setUploading(true);

    try {
      const formData = new FormData();
      formData.append('file', file);
      const response = await fetch('http://localhost:8080/api/v1/upload', {
        method: 'POST',
        body: formData,
      });
      if (response.ok) {
        console.log("File uploaded")
      } else {
        alert('File upload failed.');
      }
    } catch (error) {
      // Handle network error or other issues
      alert('An error occurred during the upload.');
    } finally {
      setUploading(false);
    }
  };

  return (
    <div className={buttonStyles.fileUploadContainer}>
      <label className={buttonStyles.labelContainer}>
        <input type="file" accept=".pdf, .doc, .docx" onChange={handleFileChange} className={buttonStyles.fileInput} />
        Choose File
      </label>
      <button name="file" onClick={handleUpload} disabled={uploading} className={buttonStyles.uploadButton}>
        {uploading ? 'Uploading...' : 'Upload File'}
      </button>
    </div>
  );
}
