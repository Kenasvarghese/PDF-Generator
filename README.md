
# PDF Generator

A simple and efficient PDF generation app that allows users to generate PDFs dynamically using predefined templates. The app exposes a POST endpoint that accepts a request body with the necessary data and configuration to generate the PDF. The template used for the PDF is customizable within the app.

## Run Locally

Clone the project

```bash
  git clone https://github.com/Kenasvarghese/PDF-Generator.git
```

Go to the project directory

```bash
  cd PDF-Generator
```

Start the server

```bash
  go run main.go
```
<em><strong>NOTE:</strong> You can change host, browserURL, template of the service by editting the corresponding file </em>

##### - Change configuration file
<em>Update values of file **./config/config.json**</em>
```sh
{
    "port":"8008",
    "basepath":"/v1",
    "browserurl":""
}
```
##### - Change PDF template 
<em>Update html of file **./pdfgenerator/template.go**</em>

##### - To get browser WS address

- Close all instances of chrome,
- Run chrome in remote debugging mode

    * On Windows:
    ```bash
    chrome.exe --remote-debugging-port=9222
    ```
    * On macOS/Linux:
    ```bash
    google-chrome --remote-debugging-port=9222
    ```
- Navigate to:
    [Remote Debugging Interface](http://localhost:9222/json)

##### - To generate PDF
 
- Send a post request
    - Path: /generate
    - Method: POST
    - Body:
        ```sh
        {
            "name":"kenas",
            "email":"name@example.com",
            "dob":"01/12/2025",
            "gender":"male"
        }
        ```
- The PDF will be generated in output.pdf        