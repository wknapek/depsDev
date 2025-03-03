import { useState } from "react";


function InsertPackage() {
  const [name, setName] = useState<string>("");
  const [json, setJson] = useState<string>("");
  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const fetchData = async () => {
      var url = new URL("http://localhost:8080"),
        params = { name: name };
      url.search = new URLSearchParams(params).toString();
      const result = await fetch(url, {
        method: "POST",
        body: json,
      });
      if (!result.ok) {
        throw new Error(`Response status: ${result.status}`);
      }
      return result.json();
    };
    fetchData();
    console.log(name);
    console.log(json);
    window.location.reload();
  };
  return (
    <div>
      <h1>Insert Package</h1>
      <form onSubmit={handleSubmit}>
        <label>Package name:</label>
        <input
          type="text"
          value={name}
          onChange={(e) => setName(e.target.value)}
          id="nameID"
        ></input>
        <br />
        <label>Package Json:</label>
        <br />
        <textarea
          cols={200}
          rows={10}
          name={json}
          id="jsonID"
          onChange={(e) => setJson(e.target.value)}
        ></textarea>
        <br />
        <button>Insert</button>
      </form>
    </div>
  );
}

export default InsertPackage;
