"use client"

import { useState } from "react";
import { Button } from "@/components/ui/button";


export default function Home() {


	const [message, setMessage] = useState<string>("");

	const fetchHello = async () => {

		try {

			const response = await fetch ("http://3.90.78.97:8080/api/hello");
			
			const data = await response.json();

			setMessage(data.message);
		
		}catch (error) {

			console.error("Error Fetching data: ", error);
		}
	
	};


  return (

	  <div className = "flex flex-col items-center justify-center h-screen">

	  <h1 className = "text-2xl font-bold mb-4"> Hello API </h1>

	  <Button onClick = {fetchHello} > Fetch Message  </Button>

	  {message && <p className="mt-4"> {message} </p>}

	  </div>

  );
}
