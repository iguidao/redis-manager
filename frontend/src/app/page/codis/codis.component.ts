import { Component, OnInit } from '@angular/core';
import { CodisService } from '../../service/codis.service';
import { CodisInfo, CodisList } from '../../service/model';
import { MatDialog } from "@angular/material/dialog";
import { DcodisComponent } from "../dialog/dcodis/dcodis.component";


@Component({
  selector: 'app-codis',
  templateUrl: './codis.component.html',
  styleUrls: ['./codis.component.css']
})
export class CodisComponent implements OnInit {

  constructor(    // private router: Router,
    public dialog: MatDialog,
    private CodisService: CodisService,
  ) { }

  error!: string;
  codislist: CodisList[] = [];
  codistotalCount!: number
  displayedColumns = ['ID', 'Cname', 'Curl','CreatedAt', 'management']

  ngOnInit() {
    this.refreshCodisInfo();
    // .subscribe(tasks => this.hist = tasks.data.lists)
  }

  refreshCodisInfo() {
    this.codislist = [];
    this.CodisService.listCodis()
    .subscribe(
      (val): void => {
          let codisinfo = val as CodisInfo;
          if (codisinfo && codisinfo.data && codisinfo.errorCode == 0) {
              this.codislist = codisinfo.data.lists
              this.codistotalCount = codisinfo.data.total
          } else{
            this.error = codisinfo.msg
          }
      }
    );
  }
  AddCodis() {
    this.dialog.open(DcodisComponent, {
      // height: '400px',
      width: '400px',
      data: {},
    });
  }
}
