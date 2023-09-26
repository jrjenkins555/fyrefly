import { useState } from 'react';

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

      const response = await fetch('http://localhost:8080/api/v1/test', {
        method: 'POST',
        body: formData,
      });

      if (response.ok) {
        // Handle success
        alert('File uploaded successfully!');
      } else {
        // Handle error
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
    <div>
      <input type="file" accept=".pdf, .doc, .docx" onChange={handleFileChange} />
      <button onClick={handleUpload} disabled={uploading}>
        {uploading ? 'Uploading...' : 'Upload File'}
      </button>
    </div>
  );
}
