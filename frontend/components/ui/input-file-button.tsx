import React, { useState } from 'react';
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Button } from "@/components/ui/button"

export default function InputFileButton() {
  const [file, setFile] = useState<null | File>(null);
  const [uploading, setUploading] = useState(false);

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files != null) {
        const selectedFile = event.target.files[0];
        setFile(selectedFile);
    }
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
    <div className="grid w-full max-w-sm items-center gap-1.5">
      <Label htmlFor="file">File</Label>
      <Input id="file" accept=".pdf, .doc, .docx" onChange={handleFileChange} type="file" />
      <Button type="submit" onClick={handleUpload} disabled={uploading}>
        {uploading ? 'Uploading...' : 'Upload File'}
      </Button>
    </div>
  )
}
