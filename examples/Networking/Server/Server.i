
function handler(||client) {
	var ip = client.ip
	print("Connection from ", ip)
	loop {
!		var message = client()
		issues {
			close(client)
			return
		}
		print(message)
	}
}


software {
	load("tcp://8000")
	issues {
		print("Could not open port 8000!")
		exit
	}
	loop {
		fork handler(open("tcp://8000"))
		issues {
			print("THERE WAS AN ERROR")
			exit
		}
	}
}
