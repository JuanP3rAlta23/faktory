// Generated by ego.
// DO NOT EDIT

//line busy.ego:1

package webui

import "fmt"
import "html"
import "io"
import "context"

import (
	"net/http"

	"github.com/contribsys/faktory/manager"
	"github.com/contribsys/faktory/server"
)

func ego_busy(w io.Writer, req *http.Request) {

//line busy.ego:13
	_, _ = io.WriteString(w, "\n\n")
//line busy.ego:14
	ego_layout(w, req, func() {
//line busy.ego:15
		_, _ = io.WriteString(w, "\n\n\n<div class=\"row header mt-3\">\n  <div class=\"col-5\">\n    <h3>")
//line busy.ego:19
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Processes"))))
//line busy.ego:19
		_, _ = io.WriteString(w, "</h3>\n  </div>\n  <div class=\"col-7\">\n    <form method=\"POST\" action=\"")
//line busy.ego:22
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(relative(req, "/busy"))))
//line busy.ego:22
		_, _ = io.WriteString(w, "\" class=\"warning-messages d-flex justify-content-end\">\n      ")
//line busy.ego:23
		_, _ = fmt.Fprint(w, csrfTag(req))
//line busy.ego:24
		_, _ = io.WriteString(w, "\n      <input type=\"hidden\" name=\"wid\" value=\"all\"/>\n      <button class=\"btn btn-primary\" type=\"submit\" name=\"signal\" value=\"quiet\" data-confirm=\"")
//line busy.ego:25
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "AreYouSure"))))
//line busy.ego:25
		_, _ = io.WriteString(w, "\">")
//line busy.ego:25
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "QuietAll"))))
//line busy.ego:25
		_, _ = io.WriteString(w, "</button>\n      <button class=\"btn btn-danger\" type=\"submit\" name=\"signal\" value=\"terminate\" data-confirm=\"")
//line busy.ego:26
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "AreYouSure"))))
//line busy.ego:26
		_, _ = io.WriteString(w, "\">")
//line busy.ego:26
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "StopAll"))))
//line busy.ego:26
		_, _ = io.WriteString(w, "</button>\n    </form>\n  </div>\n</div>\n\n<div class=\"table-responsive\">\n  <table class=\"processes table table-hover table-bordered table-striped table-light\">\n    <thead>\n      <th>")
//line busy.ego:34
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "ID"))))
//line busy.ego:34
		_, _ = io.WriteString(w, "</th>\n      <th>")
//line busy.ego:35
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Name"))))
//line busy.ego:35
		_, _ = io.WriteString(w, "</th>\n      <th>")
//line busy.ego:36
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Started"))))
//line busy.ego:36
		_, _ = io.WriteString(w, "</th>\n      <th>")
//line busy.ego:37
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Connections"))))
//line busy.ego:37
		_, _ = io.WriteString(w, "</th>\n      <th>")
//line busy.ego:38
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "RSS"))))
//line busy.ego:38
		_, _ = io.WriteString(w, "</th>\n      <th>")
//line busy.ego:39
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Busy"))))
//line busy.ego:39
		_, _ = io.WriteString(w, "</th>\n      <th>&nbsp;</th>\n    </thead>\n    ")
//line busy.ego:42
		busyWorkers(req, func(worker *server.ClientData) {
//line busy.ego:43
			_, _ = io.WriteString(w, "\n      <tr>\n        <td>\n          <code>\n            ")
//line busy.ego:46
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(worker.Wid)))
//line busy.ego:47
			_, _ = io.WriteString(w, "\n          </code>\n        </td>\n        <td>\n          <code>")
//line busy.ego:50
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(worker.Hostname)))
//line busy.ego:50
			_, _ = io.WriteString(w, ":")
//line busy.ego:50
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(worker.Pid)))
//line busy.ego:50
			_, _ = io.WriteString(w, "</code>\n          ")
//line busy.ego:51
			for _, label := range worker.Labels {
//line busy.ego:52
				_, _ = io.WriteString(w, "\n            <span class=\"badge bg-primary\">")
//line busy.ego:52
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(label)))
//line busy.ego:52
				_, _ = io.WriteString(w, "</span>\n          ")
//line busy.ego:53
			}
//line busy.ego:54
			_, _ = io.WriteString(w, "\n          ")
//line busy.ego:54
			if worker.IsQuiet() {
//line busy.ego:55
				_, _ = io.WriteString(w, "\n            <span class=\"badge bg-danger\">quiet</span>\n          ")
//line busy.ego:56
			}
//line busy.ego:57
			_, _ = io.WriteString(w, "\n        </td>\n        <td>")
//line busy.ego:58
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(Timeago(worker.StartedAt))))
//line busy.ego:58
			_, _ = io.WriteString(w, "</td>\n        <td>")
//line busy.ego:59
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(worker.ConnectionCount())))
//line busy.ego:59
			_, _ = io.WriteString(w, "</td>\n        <td>")
//line busy.ego:60
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(displayRss(worker.RssKb))))
//line busy.ego:60
			_, _ = io.WriteString(w, "</td>\n        <td>")
//line busy.ego:61
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(ctx(req).Server().Manager().BusyCount(worker.Wid))))
//line busy.ego:61
			_, _ = io.WriteString(w, "</td>\n        <td>\n          <div class=\"btn-group d-flex justify-content-end\">\n            <form method=\"POST\">\n              ")
//line busy.ego:65
			_, _ = fmt.Fprint(w, csrfTag(req))
//line busy.ego:66
			_, _ = io.WriteString(w, "\n              <input type=\"hidden\" name=\"wid\" value=\"")
//line busy.ego:66
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(worker.Wid)))
//line busy.ego:66
			_, _ = io.WriteString(w, "\"/>\n              <div class=\"text-end\">\n                ")
//line busy.ego:68
			if !worker.IsQuiet() {
//line busy.ego:69
				_, _ = io.WriteString(w, "\n                  <button class=\"btn btn-primary btn-sm\" type=\"submit\" name=\"signal\" value=\"quiet\">")
//line busy.ego:69
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Quiet"))))
//line busy.ego:69
				_, _ = io.WriteString(w, "</button>\n                ")
//line busy.ego:70
			}
//line busy.ego:71
			_, _ = io.WriteString(w, "\n                <button class=\"btn btn-danger btn-sm\" type=\"submit\" name=\"signal\" value=\"terminate\">")
//line busy.ego:71
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Stop"))))
//line busy.ego:71
			_, _ = io.WriteString(w, "</button>\n              </div>\n            </form>\n          </div>\n        </td>\n      </tr>\n    ")
//line busy.ego:77
		})
//line busy.ego:78
		_, _ = io.WriteString(w, "\n  </table>\n</div>\n\n<div class=\"row header mt-3\">\n  <div class=\"col-12\">\n    <h3>")
//line busy.ego:83
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Jobs"))))
//line busy.ego:83
		_, _ = io.WriteString(w, "</h3>\n  </div>\n</div>\n\n<div class=\"table-responsive\">\n  <table class=\"workers table table-hover table-bordered table-striped table-light\">\n    <thead>\n      <th>")
//line busy.ego:90
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Process"))))
//line busy.ego:90
		_, _ = io.WriteString(w, "</th>\n      <th>")
//line busy.ego:91
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "JID"))))
//line busy.ego:91
		_, _ = io.WriteString(w, "</th>\n      <th>")
//line busy.ego:92
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Queue"))))
//line busy.ego:92
		_, _ = io.WriteString(w, "</th>\n      <th>")
//line busy.ego:93
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Job"))))
//line busy.ego:93
		_, _ = io.WriteString(w, "</th>\n      <th>")
//line busy.ego:94
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Arguments"))))
//line busy.ego:94
		_, _ = io.WriteString(w, "</th>\n      <th>")
//line busy.ego:95
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Started"))))
//line busy.ego:95
		_, _ = io.WriteString(w, "</th>\n    </thead>\n    ")
//line busy.ego:97
		busyReservations(req, func(res *manager.Reservation) {
//line busy.ego:98
			_, _ = io.WriteString(w, "\n      ")
//line busy.ego:98
			job := res.Job
//line busy.ego:99
			_, _ = io.WriteString(w, "\n      <tr>\n        <td>\n          <code>\n            ")
//line busy.ego:102
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(res.Wid)))
//line busy.ego:103
			_, _ = io.WriteString(w, "\n          </code>\n        </td>\n        <td>\n          <code>\n            ")
//line busy.ego:107
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(job.Jid)))
//line busy.ego:108
			_, _ = io.WriteString(w, "\n          </code>\n        </td>\n        <td>\n          <a href=\"")
//line busy.ego:111
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(root(req))))
//line busy.ego:111
			_, _ = io.WriteString(w, "/queues/")
//line busy.ego:111
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(job.Queue)))
//line busy.ego:111
			_, _ = io.WriteString(w, "\">")
//line busy.ego:111
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(job.Queue)))
//line busy.ego:111
			_, _ = io.WriteString(w, "</a>\n        </td>\n        <td><code>")
//line busy.ego:113
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(job.Type)))
//line busy.ego:113
			_, _ = io.WriteString(w, "</code></td>\n        <td>\n          <div class=\"args\">")
//line busy.ego:115
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(displayArgs(job.Args))))
//line busy.ego:115
			_, _ = io.WriteString(w, "</div>\n        </td>\n        <td>")
//line busy.ego:117
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(relativeTime(res.Since))))
//line busy.ego:117
			_, _ = io.WriteString(w, "</td>\n      </tr>\n    ")
//line busy.ego:119
		})
//line busy.ego:120
		_, _ = io.WriteString(w, "\n  </table>\n</div>\n")
//line busy.ego:122
	})
//line busy.ego:123
	_, _ = io.WriteString(w, "\n")
//line busy.ego:123
}

var _ fmt.Stringer
var _ io.Reader
var _ context.Context
var _ = html.EscapeString
