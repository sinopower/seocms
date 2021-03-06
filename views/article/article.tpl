<input type="hidden" name="location" value="{{.Category.NameEn}}" />
{{with .Article}}
<h1>{{.Title}}</h1>
<div class="article-meta">
    <p>{{dateformat .Pubdate "2006-01-02"}}</p>
    {{$tags := findTags .Id}}
    <ul>
        <li>关键字：</li>
        {{str2html $tags}}
    </ul>
    <div style="clear: both;"></div>
</div><!-- End .article-meta -->
<div class="article-content">
    {{str2html .ContentHtml}}
</div><!-- End .article-content -->
{{end}}
