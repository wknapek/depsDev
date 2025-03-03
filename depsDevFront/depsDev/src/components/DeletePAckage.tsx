import { useState } from "react";

function DeletePackage() {
  const [name, setName] = useState<string>("");
  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    console.log(name)
    const fetchData = async () => {
      var url = new URL("http://localhost:8080"),
        params = { name: name };
      url.search = new URLSearchParams(params).toString();
      const result = await fetch(url, {
        method: "DELETE",
        headers:{'Access-Control-Allow-Origin': '*','Access-Control-Allow-Headers': 'Origin, X-Requested-With, Content-Type, Accept, Z-Key','Access-Control-Allow-Methods': 'GET, POST, PUT, DELETE, OPTIONS'},
      });
      if (!result.ok) {
        throw new Error(`Response status: ${result.status}`);
      }
      return result.json();
    };
    fetchData();
    console.log(name);
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
        <button>Delete</button>
      </form>
    </div>
  );
}

export default DeletePackage;