<form id="myform" method="post" action="search.php?id=search3">
       <div class="col1"> 
   <label>Ονομασία Γραμμής:  <input type="text" name="rname" pattern="[a-zA-Z0-9\u0374-\u03FF\.\-\!\?\; ]{0,}"></label>
   <label>Αριθμός Γραμμής:<input type="text" name="rnum" pattern="[a-zA-Z0-9\u0374-\u03FF\.\-\!\?\; ]{0,}"></label>
   <label>Ονομασία Στάσης:<input type="text" name="sname" pattern="[a-zA-Z0-9\u0374-\u03FF\.\-\!\?\; ]{0,}"></label>
   </div> 

  <div class="clear"></div>
 <label><div style="text-align: center;">	<button type="submit" class="buttonok"> Αναζήτηση </button></div></label>
 
 
 </form>

<?php
