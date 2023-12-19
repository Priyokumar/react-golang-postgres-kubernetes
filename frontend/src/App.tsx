import { useState } from 'react';
import './App.css';
import axios from 'axios';


function App() {

  const [data, setData] = useState<any>();
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [error, setError] = useState<string>('');
  // const url = "http://localhost:8000";
  const url = "http://34.42.54.62";

  const getData = async () => {
    setIsLoading(true);
    try {
      const resp = await axios.get(url)
      setData(resp.data);
    } catch (error: any) {
      setError(error)
    }

  }

  return (
    <div className="App">
      <header>
        <div>My App</div>
        <div>
          <span>Hello world</span>
        </div>
      </header>
      <section>
        <button onClick={() => { getData() }}>Get data</button>
        {data && (
          <div>{data.message}</div>
        )}
        {error && (
          <div>{error}</div>
        )}
      </section>
    </div>
  );
}

export default App;
