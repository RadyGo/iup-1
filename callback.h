#ifndef __GO_IUP_H 
#define __GO_IUP_H

/*
	Copyright (C) 2011 by Jeremy Cowgar <jeremy@cowgar.com>

	This file is part of go-iup.

	go-iup is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as
	published by the Free Software Foundation, either version 3 of
	the License, or (at your option) any later version.

	go-iup is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU Lesser General Public
	License along with go-iup.  If not, see <http://www.gnu.org/licenses/>.
*/


#include <iup.h>

// IUP definitions not defined
#define IUP_UNMAP_CB        "UNMAP_CB"
#define IUP_DESTROY_CB      "DESTROY_CB"
#define IUP_CARET_CB        "CARET_CB"
#define IUP_DBLCLICK_CB     "DBLCLICK_CB"
#define IUP_EDIT_CB         "EDIT_CB"
#define IUP_MULTISELECT_CB  "MULTISELECT_CB"
#define IUP_VALUECHANGED_CB "VALUECHANGED_CB"
#define IUP_TABCHANGE_CB    "TABCHANGE_CB"
#define IUP_TABCHANGEPOS_CB "TABCHANGEPOS_CB"
#define IUP_SPIN_CB         "SPIN_CB"
#define IUP_TRAYCLICK_CB    "TRAYCLICK_CB"

#define GO_PREFIX "_GO_"

extern const char *GO_MAP_CB;
extern const char *GO_UNMAP_CB;
extern const char *GO_DESTROY_CB;
extern const char *GO_GETFOCUS_CB;
extern const char *GO_KILLFOCUS_CB;
extern const char *GO_ENTERWINDOW_CB;
extern const char *GO_LEAVEWINDOW_CB;
extern const char *GO_K_ANY_CB;
extern const char *GO_HELP_CB;
extern const char *GO_BUTTON_CB;
extern const char *GO_DROPFILES_CB;
extern const char *GO_ACTION;
extern const char *GO_LIST_ACTION;
extern const char *GO_CARET_CB;
extern const char *GO_DBLCLICK_CB;
extern const char *GO_EDIT_CB;
extern const char *GO_MOTION_CB;
extern const char *GO_MULTISELECT_CB;
extern const char *GO_VALUECHANGED_CB;
extern const char *GO_TABCHANGE_CB;
extern const char *GO_TABCHANGEPOS_CB;
extern const char *GO_SPIN_CB;
extern const char *GO_SHOW_CB;
extern const char *GO_TRAYCLICK_CB;
extern const char *GO_IDLE_ACTION;


extern void goIupSetMapFunc(Ihandle *ih, int f);
extern void goIupSetUnmapFunc(Ihandle *ih, int f);
extern void goIupSetDestroyFunc(Ihandle *ih, int f);
extern void goIupSetGetFocusFunc(Ihandle *ih, int f);
extern void goIupSetKillFocusFunc(Ihandle *ih, int f);
extern void goIupSetEnterWindowFunc(Ihandle *ih, int f);
extern void goIupSetLeaveWindowFunc(Ihandle *ih, int f);
extern void goIupSetKAnyFunc(Ihandle *ih, int f);
extern void goIupSetHelpFunc(Ihandle *ih, int f);
extern void goIupSetButtonFunc(Ihandle *ih, int f);
extern void goIupSetDropFilesFunc(Ihandle *ih, int f);
extern void goIupSetActionFunc(Ihandle *ih, int f);
extern void goSetFunc(Ihandle *ih, char *goName, void *gof, char *cName, void *cf);
extern void goIupSetListActionFunc(Ihandle *ih, int f);
extern void goIupSetCaretFunc(Ihandle *ih, int f);
extern void goIupSetDblclickFunc(Ihandle *ih, int  f);
extern void goIupSetEditFunc(Ihandle *ih, int f);
extern void goIupSetMotionFunc(Ihandle *ih, int f);
extern void goIupSetMultiselectFunc(Ihandle *ih, int f);
extern void goIupSetValueChangedFunc(Ihandle *ih, int f);
extern void goIupSetTextActionFunc(Ihandle *ih, int f);
extern void goIupSetToggleActionFunc(Ihandle *ih, int f);
extern void goIupSetTabChangeFunc(Ihandle *ih, int f);
extern void goIupSetTabChangePosFunc(Ihandle *ih, int f);
extern void goIupSetSpinFunc(Ihandle *ih, int f);
extern void goIupSetShowFunc(Ihandle *ih, int f);
extern void goIupSetTrayClickFunc(Ihandle *ih, int f);
extern void goIupSetIdleFunc();
#endif

