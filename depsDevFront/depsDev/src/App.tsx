import { BrowserRouter, Routes, Route,Link } from "react-router-dom"
import ListPackages from "./components/Packages"
import InsertPackage from "./components/CreatePakage"
import UpdatePackage from "./components/UpdatePackage"
import DeletePackage from "./components/DeletePAckage"

function App(){
  return <div className="App">
    <h5>Package CRUD operations</h5>
    <BrowserRouter>
    <nav>
      <ul>
        <li>
          <Link to="/">Package list</Link>
        </li>
        <li>
          <Link to="/create">Packade Add</Link>
        </li>
        <li>
          <Link to="/delete">Delete Package</Link>
        </li>
      </ul>
    </nav>
    <Routes>
      <Route index element={<ListPackages/>} />
      <Route path="/create" element={<InsertPackage></InsertPackage>} />
      <Route path="/:name/updae" element={<UpdatePackage></UpdatePackage>} />
      <Route path="/delete" element={<DeletePackage></DeletePackage>} />
    </Routes>
    </BrowserRouter>
  </div>
}

export default App