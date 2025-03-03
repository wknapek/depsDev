import React from "react";
import { useEffect } from "react";
import styled from "styled-components";
interface apiProps{
    name:string 
	documentation:Document;    
	score:number;         
	reason:string;       
	details:string[];     
	overallScore:number;       
}

interface Document{
    description:string;
    url:string;
}

function ListPackages(){
    const [apiData,serApiData]=React.useState([]);
    useEffect(()=>{
        getUsers();
    }, []);
    const getUsers = async()=>{
        const response = await fetch('http://localhost:8080')
        const data = await response.json()
        console.log(data)
        serApiData(data)
    }
    return(
        <div>
            <h1>Package list</h1>
            {apiData.length > 0 ? apiData.map((dataPkg:apiProps)=><Props data={dataPkg}/>): ('loading')}
        </div>
    );
    
}

interface Props{
    data : apiProps
}

const Props:React.FC<Props> = ({data}) =>{
    return (<DataWrapper>
        <table>
            <thead>
                <tr>
                    <th>name</th>
                    <th>details</th>
                    <th>documentation.description</th>
                    <th>documentation.url</th>
                    <th>score</th>
                    <th>reason</th>
                    <th>details</th>
                    <th>overallScore</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td>{data.name}</td>
                    <td>{data.details}</td>
                    <td>{data.documentation.description}</td>
                    <td>{data.documentation.url}</td>
                    <td>{data.score}</td>
                    <td>{data.reason}</td>
                    <td>{data.details}</td>
                    <td>{data.overallScore}</td>
                </tr>
            </tbody>
        </table>
    </DataWrapper>
    )
}

export default ListPackages

const DataWrapper = styled.div`
>h1{
margin-top: 2rem;
text-align: center;
}
table{
border-collapse: collapse;
margin: auto
}
th{
background-color: #f2f2f2;
border: 1px solid red;
padding: 2px;
width: 10%
}
td{
border: 1px solid #dddddd;
text-align: center;
padding: 2px;
width: auto
}
`;