import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpRequest} from '@angular/common/http';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class KodoService {
  headers: Headers; 

  constructor(private http: HttpClient) { 
    this.headers = new Headers(); 
    this.headers.set('Content-Type', 'multipart/form-data'); 
  }
  private base = environment.UploadKodoUrl;

  changeListener(kodotoken: string, $event): void { 
    this.postFile(kodotoken, $event.target); 
  } 

  postFile(kodotoken: string, fileinfo: any): void { 

    var formData = new FormData(); 
    formData.append("key", "test.log"); 
    formData.append("file", fileinfo.files[0]); 

    this.http.post(this.base, formData); 
  }


  UploadImages(kodotoken: string, fileinfo: any){
    const formData: FormData = new FormData();

    formData.append('file', fileinfo);
    formData.append('key', 'ab.log');
    formData.append('token', kodotoken);
    const req = new HttpRequest('POST', 'https://upload-z1.qiniup.com', formData, {
      reportProgress: true,
      responseType: 'json'
    });

    return this.http.request(req);
    // console.log("daowole")
    // return this.http.post(this.base, { token: kodotoken, key: "test.log", file: fileinfo})
  }


//   UploadAdapter() {
//     constructor(loader) {
//         this.loader = loader;
//     }
//     upload() {
//         return new Promise((resolve, reject) => {
//         const data = new FormData();
//         data.append('upload', this.loader.file);
//         data.append('allowSize', 10);//允许图片上传的大小/兆
//         $.ajax({
//             url: '/QingXiao/Article/uploadArticleImage4',
//             type: 'POST',
//             data: data,
//             dataType: 'json',
//             processData: false,
//             contentType: false,
//             success: function (data) {
//                 if (data.res) {
//                     resolve({
//                         default: data.url
//                     });
//                 } else {
//                     reject(data.msg);
//                 }
  
//             }
//         });
  
//     });
//     }
//     abort() {
//     }
//   }

}
